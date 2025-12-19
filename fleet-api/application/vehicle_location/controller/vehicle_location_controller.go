package controller

import (
	"fleet-api/application/vehicle_location/service"
	"fleet-api/config"
	"fleet-api/domain"
	"fleet-api/helpers/errorhandler"
	"fleet-api/helpers/response"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

type vehicleLocationController struct {
	serv service.VehicleLocationServiceContract
}

func NewVehicleLocationController(serv service.VehicleLocationServiceContract) *vehicleLocationController {
	return &vehicleLocationController{
		serv: serv,
	}
}

func (c *vehicleLocationController) InitPoolData() *config.Config {
	poolData := config.DataPool.Get().(*config.Config)
	return poolData
}

func (c *vehicleLocationController) GetLast(ctx *fiber.Ctx) error {
	poolData := c.InitPoolData()

	vehicleID := ctx.Params("vehicleID")

	data, err := c.serv.GetLast(ctx.Context(), poolData, vehicleID)
	if err.Code != 0 {
		response.WriteResponse(ctx, response.Response{}, err, err.HTTPCode)
		return nil
	}

	resp := response.ResponseSuccess(data, "Success Menampilkan Data")
	response.WriteResponse(ctx, resp, domain.ErrorData{}, fasthttp.StatusOK)
	return nil
}

func (c *vehicleLocationController) GetHistory(ctx *fiber.Ctx) error {
	poolData := c.InitPoolData()

	vehicleID := ctx.Params("vehicleID")
	start, errData := strconv.ParseInt(ctx.Query("start"), 10, 64)
	if errData != nil {
		err := errorhandler.ErrInternal(errorhandler.ErrCodeParsing, errData)
		response.WriteResponse(ctx, response.Response{}, err, err.HTTPCode)
		return nil
	}

	end, errData := strconv.ParseInt(ctx.Query("end"), 10, 64)
	if errData != nil {
		err := errorhandler.ErrInternal(errorhandler.ErrCodeParsing, errData)
		response.WriteResponse(ctx, response.Response{}, err, err.HTTPCode)
		return nil
	}

	if start > end {
		err := errorhandler.ErrValidation(fmt.Errorf(fmt.Sprintf("%s : start cannot be bigger than end", errorhandler.ErrInvalidDataQuery)))
		response.WriteResponse(ctx, response.Response{}, err, err.HTTPCode)
		return nil
	}

	data, err := c.serv.GetHistory(ctx.Context(), poolData, vehicleID, start, end)
	if err.Code != 0 {
		response.WriteResponse(ctx, response.Response{}, err, err.HTTPCode)
		return nil
	}

	resp := response.ResponseSuccess(data, "Success Menampilkan Data")
	response.WriteResponse(ctx, resp, domain.ErrorData{}, fasthttp.StatusOK)
	return nil
}
