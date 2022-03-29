package server

// import (
// 	"testing"

// 	"github.com/sonr-io/highway-go/config"
// 	"github.com/sonr-io/highway-go/models"
// 	"github.com/stretchr/testify/suite"
// )

// type ServerSuite struct {
// 	config *config.Config
// 	server *Server

// 	suite.Suite
// }

// func (ss *ServerSuite) SetupSuite() {
// 	ss.config = &config.Config{
// 		DBName:      "sqlite3",
// 		DBPath:      ":memory:",
// 		HostAddress: "localhost",
// 	}
// 	err := models.Setup(ss.config)
// 	ss.Nil(err)

// 	ss.server, err = NewServer(nil, ss.config)
// 	ss.Nil(err)
// }

// func TestRunServerSuite(t *testing.T) {
// 	suite.Run(t, new(ServerSuite))
// }
