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
	surveyDocType            service.SurveyDocTypeService
	typeOfSurvey             service.TypeOfSurveyService
	province                 service.ProvinceService
}

// NewHandler สร้าง Handler ใหม่
func NewHandler() *Handler {
	userService := service.NewUserService()
	landOfficeService := service.NewLandOfficeService()
	getSvaCadastral01Service := service.NewGetSVACadastral01Service()
	getSVACadastralImageNull := service.NewGet02getSVACadastralImageNullService()
	summarySheetCode := service.NewSummarySheetCodeService()
	surveyDocType := service.NewSurveyDocTypeService()
	typeOfSurvey := service.NewTypeOfSurveyService()
	province := service.NewProvinceService()
	return &Handler{userService: *userService, landOfficeService: *landOfficeService, getSvaCadastral01Service: *getSvaCadastral01Service, getSVACadastralImageNull: *getSVACadastralImageNull, summarySheetCode: *summarySheetCode, surveyDocType: *surveyDocType, typeOfSurvey: *typeOfSurvey, province: *province} // ส่งค่าไปยัง Handler
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

func (h *Handler) GetSummaryBoxBySheetCode(c *fiber.Ctx) error {
	var request struct {
		LANDOFFICE_SEQ int64 `json:"LANDOFFICE_SEQ"` // Expecting cadastralSeq in the body as int64
		SHEETCODE      int64 `json:"SHEETCODE"`      // Expecting cadastralSeq in the body as int64
	}

	if err := c.BodyParser(&request); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}
	summaryBox, err := h.summarySheetCode.FetchSummaryBoxBySheetCode(request.LANDOFFICE_SEQ, request.SHEETCODE) // เรียกใช้ Service
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching 01getSVACadastral")
	}
	return c.JSON(summaryBox)
}

func (h *Handler) GetCadastralKeyin(c *fiber.Ctx) error {
	var request struct {
		LANDOFFICE_SEQ int64  `json:"LANDOFFICE_SEQ"` // Expecting cadastralSeq in the body as int64
		SHEETCODE      int64  `json:"SHEETCODE"`      // Expecting cadastralSeq in the body as int64
		BOX_NO         string `json:"BOX_NO"`         // Expecting cadastralSeq in the body as int64
	}

	if err := c.BodyParser(&request); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}
	cadastralKeyin, err := h.summarySheetCode.FetchGetCadastralKeyin(request.LANDOFFICE_SEQ, request.SHEETCODE, request.BOX_NO) // เรียกใช้ Service
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching GetCadastralKeyin")
	}

	return c.JSON(cadastralKeyin)
}

func (h *Handler) GetSurveyDocTypeGroup(c *fiber.Ctx) error {
	surveyDocTypeGroups, err := h.surveyDocType.FetchSurveyDocType()
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching SurveyDocTypeGroup")
	}
	return c.JSON(surveyDocTypeGroups)
}

func (h *Handler) GetListAllKeyin(c *fiber.Ctx) error {
	var request struct {
		LANDOFFICE_SEQ  int64  `json:"LANDOFFICE_SEQ"`  // Expecting cadastralSeq in the body as int64
		PAGE            int64  `json:"PAGE"`            // Expecting cadastralSeq in the body as int64
		PAGE_ROW        int64  `json:"PAGE_ROW"`        // Expecting cadastralSeq in the body as int64
		START_SHEETCODE *int64 `json:"START_SHEETCODE"` // Expecting cadastralSeq in the body as int64
		END_SHEETCODE   *int64 `json:"END_SHEETCODE"`   // Expecting cadastralSeq in the body as int64
		START_BOX       *int64 `json:"START_BOX"`       // Expecting cadastralSeq in the body as int64
		END_BOX         *int64 `json:"END_BOX"`         // Expecting cadastralSeq in the body as int64
	}

	if err := c.BodyParser(&request); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}
	listAllKeyin, err := h.summarySheetCode.FetchGetListAllKeyin(request.LANDOFFICE_SEQ, request.PAGE, request.PAGE_ROW, request.START_SHEETCODE, request.END_SHEETCODE, request.START_BOX, request.END_BOX) // เรียกใช้ Service
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching GetListAllKeyin")
	}

	return c.JSON(listAllKeyin)
}

func (h *Handler) Get02CadastralImage(c *fiber.Ctx) error {
	var request struct {
		CADASTRAL_SEQ int64 `json:"CADASTRAL_SEQ"` // Expecting cadastralSeq in the body as int64
	}

	if err := c.BodyParser(&request); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}
	cadastralImage, err := h.getSVACadastralImageNull.Get02CadastralImageService(request.CADASTRAL_SEQ) // เรียกใช้ landOfficeService
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching Get02CadastralImage")
	}
	return c.JSON(cadastralImage)
}

func (h *Handler) GetTypeOfSurvey(c *fiber.Ctx) error {
	surveyDocType, err := h.typeOfSurvey.FetchTypeOfSurvey() // เรียกใช้ landOfficeService
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching Get02CadastralImage")
	}
	return c.JSON(surveyDocType)
}

func (h *Handler) GetAllProvince(c *fiber.Ctx) error {
	provinces, err := h.province.FetchAllProvinces() // เรียกใช้ landOfficeService
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching land offices")
	}
	return c.JSON(provinces)
}
