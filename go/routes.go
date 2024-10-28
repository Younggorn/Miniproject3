package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func getRoomHandler(c *fiber.Ctx) error {
	fmt.Println("getRoomHandler")
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	room, err := getRoom(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.SendStatus(fiber.StatusNotFound)
		}
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(room)
}

func getbuildingtype(c *fiber.Ctx) error {
	buildingtype, err := buildingtype()
	if err != nil {
		if err == sql.ErrNoRows {
			return c.SendStatus(fiber.StatusNotFound)
		}
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(buildingtype)

}

func getroomtype(c *fiber.Ctx) error {
	roomtype, err := roomtype()
	if err != nil {
		if err == sql.ErrNoRows {
			return c.SendStatus(fiber.StatusNotFound)
		}
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(roomtype)

}

func getpicture(c *fiber.Ctx) error {
	fmt.Println("getpicture")
	pic, err := getpic()
	if err != nil {
		if err == sql.ErrNoRows {
			return c.SendStatus(fiber.StatusNotFound)
		}
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(pic)

}
func getAddress_id(c *fiber.Ctx) error {
	address, err := getAddress()
	if err != nil {
		if err == sql.ErrNoRows {
			return c.SendStatus(fiber.StatusNotFound)
		}
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(address)

}
func getstatustype(c *fiber.Ctx) error {
	statustype, err := statustype()
	if err != nil {
		if err == sql.ErrNoRows {
			return c.SendStatus(fiber.StatusNotFound)
		}
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(statustype)

}
func getfloortype(c *fiber.Ctx) error {
	floortype, err := floortype()
	if err != nil {
		if err == sql.ErrNoRows {
			return c.SendStatus(fiber.StatusNotFound)
		}
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(floortype)

}
func getRoomsHandler(c *fiber.Ctx) error {
	fmt.Println("getRoomsHandler")
	rooms, err := getRooms()
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("ErrNoRows")

			return c.SendStatus(fiber.StatusNotFound)
		}
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(rooms)
}

func createRoomHandler(c *fiber.Ctx) error {
	room := new(Room)

	room.ID, _ = strconv.Atoi(c.FormValue("id"))
	room.Cap, _ = strconv.Atoi(c.FormValue("cap"))
	room.RoomTypeID, _ = strconv.Atoi(c.FormValue("room_type_id"))
	room.AddressID, _ = strconv.Atoi(c.FormValue("address_id"))

	room.Status, _ = strconv.Atoi(c.FormValue("status"))

	room.Name = c.FormValue("name")
	room.Description = c.FormValue("description")

	file, err := c.FormFile("roompic")
	if err != nil {
		fmt.Println("Error receiving file:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot receive file",
		})
	}

	fileContent, err := file.Open()
	if err != nil {
		fmt.Println("Error opening file:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot open file",
		})
	}
	defer fileContent.Close()
	roompicBytes, err := ioutil.ReadAll(fileContent)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot read file",
		})
	}

	room.Roompic = roompicBytes

	if err := createRoom(room); err != nil {
		fmt.Println("Error creating room:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot create room",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Create Room Successfully",
	})
}

func updateRoomHandler(c *fiber.Ctx) error {
	fmt.Println("updateRoomHandler")

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	room := new(Room)
	err = c.BodyParser(room)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	err = updateRoom(id, room)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(room)
}

func deleteRoomHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	err = deleteRoom(id)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"message": "Delete Room Successfully",
	})
}

func getDepartmentsHandler(c *fiber.Ctx) error {
	departments, err := getDepartments()
	if err != nil {
		if err == sql.ErrNoRows {
			return c.SendStatus(fiber.StatusNotFound)
		}
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(departments)
}

func getRolesHandler(c *fiber.Ctx) error {
	roles, err := getRoles()
	if err != nil {
		if err == sql.ErrNoRows {
			return c.SendStatus(fiber.StatusNotFound)
		}
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(roles)
}

func getMenusHandler(c *fiber.Ctx) error {
	menus, err := getMenus()
	if err != nil {
		if err == sql.ErrNoRows {
			return c.SendStatus(fiber.StatusNotFound)
		}
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(menus)
}

func getEmployeesHandler(c *fiber.Ctx) error {
	employees, err := getEmployees()
	if err != nil {
		if err == sql.ErrNoRows {
			return c.SendStatus(fiber.StatusNotFound)
		}
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(employees)
}

func getEmployeeHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	employees, err := getEmployee(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.SendStatus(fiber.StatusNotFound)
		}
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(employees)
}

func createEmlpoyeeHandler(c *fiber.Ctx) error {
	employee := new(Employee)
	err := c.BodyParser(&employee)
	if err != nil {
		return err
	}
	err = createEmployee(employee)
	if err != nil {
		return err
	}
	return c.JSON(employee)
}

func updateEmployeeHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	var employee Employee
	err = c.BodyParser(employee)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	err = updateEmployee(id, &employee)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(employee)
}

func getPermissionsHandler(c *fiber.Ctx) error {
	permissions, err := getPermissions()
	if err != nil {
		if err == sql.ErrNoRows {
			return c.SendStatus(fiber.StatusNotFound)
		}
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(permissions)
}

func getPermissionsUserHandler(c *fiber.Ctx) error {
	token := c.Locals(userContextKey).(*Auth)
	userEmail := token.Email
	permissions, err := getPermissionsUser(userEmail)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.JSON(permissions)
}

func updatePermissionsHandler(c *fiber.Ctx) error {
	var permissions []Permission
	err := c.BodyParser(&permissions)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	err = updatePermission(id, permissions)
	if err != nil {
		fmt.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(fiber.Map{
		"message": "Update Permission Successfully",
	})
}

func getRoomsBookedHandler(c *fiber.Ctx) error {
	bookings, err := getBookings()
	if err != nil {
		fmt.Println(err)
		return c.SendStatus(fiber.ErrBadGateway.Code)
	}
	return c.JSON(bookings)
}

func loginHandler(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if err := verifyUser(user.Email, user.Password); err != nil {
		return err
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["Email"] = user.Email
	claims["Exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"message": "Login Success",
		"token":   t,
	})
}

func registerHandler(c *fiber.Ctx) error {
	employee := new(Employee)
	err := c.BodyParser(&employee)
	if err != nil {
		return err
	}
	err = createEmployee(employee)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"message": "Register Successfully",
	})
}
func bookRoomHandler(c *fiber.Ctx) error {
	token := c.Locals(userContextKey).(*Auth)
	userEmail := token.Email

	var Booking Booking

	if err := c.BodyParser(&Booking); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	fmt.Println("userEmail:", userEmail)
	err := db.QueryRow(`SELECT id FROM employee WHERE email = :1`, userEmail).Scan(&Booking.EmpID)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Employee not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to query employee ID",
		})
	}
	fmt.Println(Booking.EmpID)

	return c.JSON(fiber.Map{
		"message": "Room booked successfully",
	})
}

func unlockRoomHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	err = unlockRoom(id)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"message": "Unlock Room Successfully",
	})
}

func cancelRoomHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	var cancel Cancel
	err = c.BodyParser(cancel)
	if err != nil {
		return err
	}
	err = cancelRoom(id, cancel)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"message": "Cancel Room Successfully",
	})
}

func getReportUsedCanceledHandler(c *fiber.Ctx) error {
	report, err := getReportUsedCanceled()
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(report)
}

func getReportLockedEmployeesHandler(c *fiber.Ctx) error {
	report, err := getReportLockEmployee()
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(report)
}
