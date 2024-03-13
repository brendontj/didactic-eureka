package server

import (
	"fmt"
	"github.com/brendontj/didactic-eureka/adapter/handler"
	"github.com/brendontj/didactic-eureka/adapter/messenger"
	"github.com/brendontj/didactic-eureka/adapter/repository"
	"github.com/brendontj/didactic-eureka/core/usecase"
	"github.com/brendontj/didactic-eureka/infrastructure/postgres"
	"github.com/brendontj/didactic-eureka/infrastructure/rabbitmq"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

type Server struct {
	*gin.Engine

	repository *repository.Adapter
	messenger  *messenger.Adapter

	createCustomerUc   usecase.CreateCustomer
	findAllCustomersUc usecase.FindAllCustomers
	findCustomerByIdUc usecase.FindCustomerById
	updateCustomerUc   usecase.UpdateCustomer
	deleteCustomerUc   usecase.DeleteCustomer
}

func NewServer() *Server {
	return &Server{
		Engine: gin.Default(),
	}
}

func (s *Server) Run() {
	defer s.Close()
	s.InitializeDependencies()
	s.InitializeUseCases()
	s.DefineRoutes()

	if err := s.Engine.Run(); err != nil {
		panic(err)
	}
}

func (s *Server) InitializeDependencies() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	pgDBInfra, err := postgres.NewDB(postgres.Config{
		User:   os.Getenv("POSTGRES_USER"),
		Pass:   os.Getenv("POSTGRES_PASSWORD"),
		Host:   os.Getenv("POSTGRES_HOST"),
		Port:   os.Getenv("POSTGRES_PORT"),
		DBName: os.Getenv("POSTGRES_DB"),
	})
	if err != nil {
		panic(err)
	}

	s.repository = repository.NewAdapter(pgDBInfra)
	rabbitmqClientInfra, err := rabbitmq.NewClient(rabbitmq.Config{
		User: os.Getenv("RABBITMQ_USER"),
		Pass: os.Getenv("RABBITMQ_PASSWORD"),
		Host: os.Getenv("RABBITMQ_HOST"),
		Port: os.Getenv("RABBITMQ_PORT"),
	})
	if err != nil {
		panic(err)
	}

	s.messenger = messenger.NewAdapter(rabbitmqClientInfra)
	if err = s.messenger.DeclareQueue(os.Getenv("SAVE_CUSTOMER_QUEUE")); err != nil {
		panic(err)
	}
}

func (s *Server) InitializeUseCases() {
	s.createCustomerUc = usecase.NewCreateCustomer(s.messenger, os.Getenv("SAVE_CUSTOMER_QUEUE"))
	s.findAllCustomersUc = usecase.NewFindAllCustomers(s.repository)
	s.findCustomerByIdUc = usecase.NewFindCustomerById(s.repository)
	s.updateCustomerUc = usecase.NewUpdateCustomer(s.repository)
	s.deleteCustomerUc = usecase.NewDeleteCustomer(s.repository)
}

func (s *Server) Close() {
	s.repository.Close()
	s.messenger.Close()
}

func (s *Server) DefineRoutes() {
	routerGroup := s.Group("/api/v1")

	routerGroup.GET("/customers", s.FindAllCustomers)
	routerGroup.GET("/customers/:id", s.FindCustomerById)
	routerGroup.POST("/customers", s.SaveCustomer)
	routerGroup.PUT("/customers/:id/version/:version", s.UpdateCustomer)
	routerGroup.DELETE("/customers/:id", s.DeleteCustomer)
}

func (s *Server) FindAllCustomers(c *gin.Context) {
	h := handler.NewFindAllCustomersHandler(s.findAllCustomersUc)
	h.Handle(c.Writer, c.Request)
}

func (s *Server) FindCustomerById(c *gin.Context) {
	h := handler.NewFindCustomerByIdHandler(s.findCustomerByIdUc)
	customerID := c.Param("id")
	fmt.Println(customerID)

	c.Request.Header.Set("id", customerID)
	h.Handle(c.Writer, c.Request)
}

func (s *Server) SaveCustomer(c *gin.Context) {
	h := handler.NewCreateCustomerHandler(s.createCustomerUc)
	h.Handle(c.Writer, c.Request)
}

func (s *Server) UpdateCustomer(c *gin.Context) {
	h := handler.NewUpdateCustomerHandler(s.updateCustomerUc)
	customerID := c.Param("id")
	version := c.Param("version")
	c.Request.Header.Set("id", customerID)
	c.Request.Header.Set("version", version)
	h.Handle(c.Writer, c.Request)
}

func (s *Server) DeleteCustomer(c *gin.Context) {
	h := handler.NewDeleteCustomerHandler(s.deleteCustomerUc)
	customerID := c.Param("id")
	c.Request.Header.Set("id", customerID)
	h.Handle(c.Writer, c.Request)
}
