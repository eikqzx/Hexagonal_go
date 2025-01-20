package http

import (
	"Hexagonal_go/app/internal/core/service"
	"log"

	"github.com/gofiber/fiber/v2"
)

// Handler สำหรับจัดการคำขอ HTTP
type Handler struct {
	userService              service.UserService
	landOfficeService        service.LandOfficeService
	getSvaCadastral01Service service.GetSVACadastral01Service
}

// NewHandler สร้าง Handler ใหม่
func NewHandler() *Handler {
	userService := service.NewUserService()
	landOfficeService := service.NewLandOfficeService()
	getSvaCadastral01Service := service.NewGetSVACadastral01Service()
	return &Handler{userService: *userService, landOfficeService: *landOfficeService, getSvaCadastral01Service: *getSvaCadastral01Service} // ส่งค่าไปยัง Handler
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
