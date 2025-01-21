package http

import (
	"Hexagonal_go/app/internal/core/model"
	"Hexagonal_go/app/internal/core/service"
	"log"

	"github.com/gofiber/fiber/v2"
)

// Handler สำหรับจัดการคำขอ HTTP
type Handler struct {
	userService              service.UserService
	landOfficeService        service.LandOfficeService
	getSvaCadastral01Service service.GetSVACadastral01Service
	getSVACadastralImageNull service.Get02getSVACadastralImageNullService
	summarySheetCode         service.SummarySheetCodeService
}

// NewHandler สร้าง Handler ใหม่
func NewHandler() *Handler {
	userService := service.NewUserService()
	landOfficeService := service.NewLandOfficeService()
	getSvaCadastral01Service := service.NewGetSVACadastral01Service()
	getSVACadastralImageNull := service.NewGet02getSVACadastralImageNullService()
	summarySheetCode := service.NewSummarySheetCodeService()
	return &Handler{userService: *userService, landOfficeService: *landOfficeService, getSvaCadastral01Service: *getSvaCadastral01Service, getSVACadastralImageNull: *getSVACadastralImageNull, summarySheetCode: *summarySheetCode} // ส่งค่าไปยัง Handler
}

// GetUsers ดึงข้อมูลผู้ใช้ทั้งหมด
func (h *Handler) GetUsers(c *fiber.Ctx) error {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching users")
	}
	return c.JSON(users)
}

// GetAllLandOffice ดึงข้อมูลสำนักงานที่ดินทั้งหมด
func (h *Handler) GetAllLandOffice(c *fiber.Ctx) error {
	landOffices, err := h.landOfficeService.GetAllLandOffice() // เรียกใช้ landOfficeService
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching land offices")
	}
	return c.JSON(landOffices)
}

func (h *Handler) GetSVALandOffice(c *fiber.Ctx) error {
	landOffices, err := h.landOfficeService.GetSVALandOffice() // เรียกใช้ landOfficeService
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching land offices")
	}
	return c.JSON(landOffices)
}

func (h *Handler) GetLandOfficeByLandOfficeSeq(c *fiber.Ctx) error {
	var request struct {
		LANDOFFICE_SEQ int64 `json:"LANDOFFICE_SEQ"` // Expecting landOfficeSeq in the body as int64
	}

	if err := c.BodyParser(&request); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}

	landOffice, err := h.landOfficeService.GetLandOfficeByLandOffice(request.LANDOFFICE_SEQ) // เรียกใช้ landOfficeService
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching land office")
	}
	return c.JSON(landOffice)
}

func (h *Handler) GetFetch01getSVACadastral(c *fiber.Ctx) error {
	var request struct {
		CADASTRAL_SEQ int64 `json:"CADASTRAL_SEQ"` // Expecting cadastralSeq in the body as int64
	}

	if err := c.BodyParser(&request); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}

	svaCadastral01, err := h.getSvaCadastral01Service.Fetch01getSVACadastral(request.CADASTRAL_SEQ) // เรียกใช้ landOfficeService
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching 01getSVACadastral")
	}
	return c.JSON(svaCadastral01)
}

func (h *Handler) GetFetch02SVACadastralImageNull(c *fiber.Ctx) error {
	var request struct {
		LANDOFFICE_SEQ int64 `json:"LANDOFFICE_SEQ"` // Expecting cadastralSeq in the body as int64
	}

	if err := c.BodyParser(&request); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}

	svaCadastral01, err := h.getSVACadastralImageNull.Fetch02getSVACadastralImageNullService(request.LANDOFFICE_SEQ) // เรียกใช้ landOfficeService
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching 01getSVACadastral")
	}
	return c.JSON(svaCadastral01)
}

func (h *Handler) Update02SVACadastralImageNull(c *fiber.Ctx) error {
	var cadastralImage model.CadastralImage
	if err := c.BodyParser(&cadastralImage); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request payload")
	}

	result, err := h.getSVACadastralImageNull.Update02SVACadastralImageNullService(cadastralImage)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	response := map[string]interface{}{
		"message":       "Cadastral image updated successfully",
		"rows_affected": rowsAffected,
		"succeed":       rowsAffected > 0,
	}
	return c.JSON(response)
}

func (h *Handler) GetSummarySheetCode(c *fiber.Ctx) error {
	var request struct {
		LANDOFFICE_SEQ int64 `json:"LANDOFFICE_SEQ"` // Expecting cadastralSeq in the body as int64
	}

	if err := c.BodyParser(&request); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}

	summarySheetCode, err := h.summarySheetCode.FetchSummarySheetCode(request.LANDOFFICE_SEQ) // เรียกใช้ landOfficeService
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching 01getSVACadastral")
	}
	return c.JSON(summarySheetCode)
}
