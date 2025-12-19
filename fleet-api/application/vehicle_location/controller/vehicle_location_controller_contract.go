package controller

import "github.com/gofiber/fiber/v2"

type VehicleLocationControllerContract interface {
	GetLast(ctx *fiber.Ctx) error
	GetHistory(ctx *fiber.Ctx) error
}
