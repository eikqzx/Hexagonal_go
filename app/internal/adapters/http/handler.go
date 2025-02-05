package http

import (
	"Hexagonal_go/app/internal/core/model"
	"Hexagonal_go/app/internal/core/service"
	"log"
	"strconv"

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
	amphurService            service.AmphurService
	tambolService            service.TambolService
	cadastralLandService     service.CadastralLandService
	getMap01                 service.GetMap01Service
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
	amphurService := service.NewAmphurService()
	tambolService := service.NewTambolService()
	cadastralLandService := service.NewCadastralLandService()
	getMap01 := service.NewGetMap01Service()
	return &Handler{
		userService: *userService, landOfficeService: *landOfficeService,
		getSvaCadastral01Service: *getSvaCadastral01Service, getSVACadastralImageNull: *getSVACadastralImageNull,
		summarySheetCode: *summarySheetCode, surveyDocType: *surveyDocType,
		typeOfSurvey: *typeOfSurvey, province: *province, amphurService: *amphurService,
		tambolService: *tambolService, cadastralLandService: *cadastralLandService,
		getMap01: *getMap01,
	} // ส่งค่าไปยัง Handler
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
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching Province")
	}
	return c.JSON(provinces)
}

func (h *Handler) GetAmphur(c *fiber.Ctx) error {
	amphurs, err := h.amphurService.FetchAllAmphur() // เรียกใช้ landOfficeService
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching Amphur")
	}
	return c.JSON(amphurs)
}

func (h *Handler) GetAmphurByProvinceID(c *fiber.Ctx) error {
	idRequest := c.Params("id")

	// var request struct {
	// 	PROVINCE_SEQ int64 `json:"PROVINCE_SEQ"` // Expecting provinceSeq in the body as int64
	// }

	id, err := strconv.ParseInt(idRequest, 10, 64)
	if err != nil {
		// หากแปลงไม่ได้ ให้ส่งข้อความ error กลับ
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	amphurs, err := h.amphurService.FetchAmphurByProvinceID(id) // เรียกใช้ landOfficeService
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching Amphur")
	}
	return c.JSON(amphurs)
}

func (h *Handler) GetTambol(c *fiber.Ctx) error {
	tambols, err := h.tambolService.FetchAllTambol() // เรียกใช้ tambolService
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching Tambol")
	}
	return c.JSON(tambols)
}

func (h *Handler) GetTambolByAmphurID(c *fiber.Ctx) error {
	idRequest := c.Params("id")

	id, err := strconv.ParseInt(idRequest, 10, 64)
	if err != nil {
		// หากแปลงไม่ได้ ให้ส่งข้อความ error กลับ
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	tambols, err := h.tambolService.FetchTambolByAmphurID(id) // เรียกใช้ tambolService
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching Tambol")
	}
	return c.JSON(tambols)
}

func (h *Handler) GetAllCadastralLand(c *fiber.Ctx) error {
	cadastralLand, err := h.cadastralLandService.FetchAllCadastralLand() // เรียกใช้ CadastralLand
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching land offices")
	}
	return c.JSON(cadastralLand)
}

func (h *Handler) GetCadastralLandByCadastralSeq(c *fiber.Ctx) error {
	idRequest := c.Params("id")

	id, err := strconv.ParseInt(idRequest, 10, 64)
	if err != nil {
		// หากแปลงไม่ได้ ให้ส่งข้อความ error กลับ
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	cadastralLand, err := h.cadastralLandService.CadastralLandByCadastralSeq(id) // เรียกใช้ CadastralLand
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching land offices")
	}
	return c.JSON(cadastralLand)
}

func (h *Handler) GetMap01(c *fiber.Ctx) error {
	var request struct {
		OGR_FID        *int64  `json:"OGR_FID"`
		LANDOFFICE_SEQ *int64  `json:"LANDOFFICE_SEQ"`
		UTMMAP1        *string `json:"UTMMAP1"`
		UTMMAP2        *string `json:"UTMMAP2"`
		UTMMAP3        *string `json:"UTMMAP3"`
		UTMMAP4        *string `json:"UTMMAP4"`
		UTMSCALE       *string `json:"UTMSCALE"`
		LAND_NO        *string `json:"LAND_NO"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	mapLandGISList, err := h.getMap01.FetchMapLandGIS(
		request.OGR_FID, request.LANDOFFICE_SEQ, request.UTMMAP1, request.UTMMAP2, request.UTMMAP3, request.UTMMAP4, request.UTMSCALE, request.LAND_NO,
	)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching map land GIS")
	}
	return c.JSON(mapLandGISList)
}

func (h *Handler) UpdateCadastralLand(c *fiber.Ctx) error {
	idRequest := c.Params("id")
	id, err := strconv.ParseInt(idRequest, 10, 64)
	if err != nil {
		// หากแปลงไม่ได้ ให้ส่งข้อความ error กลับ
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}
	var cadastralLand model.CadastralLand
	if err := c.BodyParser(&cadastralLand); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request payload")
	}

	result, err := h.cadastralLandService.UpdateCadastralLand(cadastralLand, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	response := map[string]interface{}{
		"message":       "Cadastral land updated successfully",
		"rows_affected": rowsAffected,
		"succeed":       rowsAffected > 0,
	}
	return c.JSON(response)
}

func (h *Handler) InsertCadastralLand(c *fiber.Ctx) error {
	var request struct {
		CadastralLandNganNum  *int64  `json:"CADASTRAL_LAND_NGAN_NUM"`
		CadastralLandRaiNum   *int64  `json:"CADASTRAL_LAND_RAI_NUM"`
		CadastralLandWaNum    *int64  `json:"CADASTRAL_LAND_WA_NUM"`
		CadastralLandSubwaNum *int64  `json:"CADASTRAL_LAND_SUBWA_NUM"`
		CadastralLandUtmap1   *string `json:"CADASTRAL_LAND_UTMAP1"`
		CadastralLandUtmap2   *string `json:"CADASTRAL_LAND_UTMAP2"`
		CadastralLandUtmap3   *string `json:"CADASTRAL_LAND_UTMAP3"`
		CadastralLandUtmap4   *string `json:"CADASTRAL_LAND_UTMAP4"`
		CadastralProvinceSeq  *int64  `json:"CADASTRAL_PROVINCE_SEQ"`
		CadastralAmphurSeq    *int64  `json:"CADASTRAL_AMPHUR_SEQ"`
		CadastralTambolSeq    *int64  `json:"CADASTRAL_TAMBOL_SEQ"`
		ScaleMap              *string `json:"SCALE_MAP"`
		ScaleSeq              *int64  `json:"SCALE_SEQ"`
		CadastralSeq          *int64  `json:"CADASTRAL_SEQ"`
		LandofficeSeq         *int64  `json:"LANDOFFICE_SEQ"`
		CreateUser            *string `json:"CREATE_USER"`
		RecordStatus          *string `json:"RECORD_STATUS"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	cadastralLand := &model.CadastralLand{
		CadastralLandNganNum:  request.CadastralLandNganNum,
		CadastralLandRaiNum:   request.CadastralLandRaiNum,
		CadastralLandWaNum:    request.CadastralLandWaNum,
		CadastralLandSubwaNum: request.CadastralLandSubwaNum,
		CadastralLandUtmap1:   request.CadastralLandUtmap1,
		CadastralLandUtmap2:   request.CadastralLandUtmap2,
		CadastralLandUtmap3:   request.CadastralLandUtmap3,
		CadastralLandUtmap4:   request.CadastralLandUtmap4,
		CadastralProvinceSeq:  request.CadastralProvinceSeq,
		CadastralAmphurSeq:    request.CadastralAmphurSeq,
		CadastralTambolSeq:    request.CadastralTambolSeq,
		ScaleMap:              request.ScaleMap,
		ScaleSeq:              request.ScaleSeq,
		CadastralSeq:          request.CadastralSeq,
		LandofficeSeq:         request.LandofficeSeq,
		CreateUser:            request.CreateUser,
		RecordStatus:          request.RecordStatus,
	}

	result, err := h.cadastralLandService.InsertCadastralLand(cadastralLand)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	response := map[string]interface{}{
		"message":       "Cadastral Land Insert successfully",
		"rows_affected": rowsAffected,
		"succeed":       rowsAffected > 0,
	}
	return c.JSON(response)
}
