package main

import (
	"github.com/aadi-1024/auth-micro/pkg/database"
	"github.com/aadi-1024/auth-micro/pkg/jwtUtil"
)

type Config struct {
	Db  *database.Database
	Jwt *jwtUtil.JwtConfig
}
