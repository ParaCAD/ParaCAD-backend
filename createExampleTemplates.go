package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/ParaCAD/ParaCAD-backend/auth"
	"github.com/ParaCAD/ParaCAD-backend/controller"
	"github.com/google/uuid"
)

func createExampleUsersAndTemplates(c *controller.Controller) {
	w := httptest.NewRecorder()
	r := createUserRequest("Tomasz", "tomasz@mail.com", "1234")
	c.HandleRegister(w, &r, nil)
	if w.Code >= 400 {
		panic(fmt.Sprintf("Expected status code 200, got %d: %s", w.Code, w.Body.String()))
	}
	userID := struct {
		UUID uuid.UUID `json:"uuid"`
	}{}
	err := json.NewDecoder(w.Body).Decode(&userID)
	if err != nil {
		panic(fmt.Sprintf("Failed to decode user ID: %v", err))
	}

	w = httptest.NewRecorder()
	r = createCubeRequest(userID.UUID.String())
	c.HandleCreateTemplate(w, &r, nil)
	if w.Code >= 400 {
		panic(fmt.Sprintf("Expected status code 200, got %d: %s", w.Code, w.Body.String()))
	}

	w = httptest.NewRecorder()
	r = createMeasuringCupRequest(userID.UUID.String())
	c.HandleCreateTemplate(w, &r, nil)
	if w.Code >= 400 {
		panic(fmt.Sprintf("Expected status code 200, got %d: %s", w.Code, w.Body.String()))
	}

	w = httptest.NewRecorder()
	r = createBoxWithLidRequest(userID.UUID.String())
	c.HandleCreateTemplate(w, &r, nil)
	if w.Code >= 400 {
		panic(fmt.Sprintf("Expected status code 200, got %d: %s", w.Code, w.Body.String()))
	}

	w = httptest.NewRecorder()
	r = createTestTubeHolderRequest(userID.UUID.String())
	c.HandleCreateTemplate(w, &r, nil)
	if w.Code >= 400 {
		panic(fmt.Sprintf("Expected status code 200, got %d: %s", w.Code, w.Body.String()))
	}
}

func createUserRequest(username string, email string, password string) http.Request {
	body := io.NopCloser(bytes.NewReader([]byte(fmt.Sprintf(`
	{
		"username": "%s",
		"email": "%s",
		"password": "%s"
	}
	`, username, email, password))))

	r := http.Request{
		Method: http.MethodPost,
		Body:   body,
	}

	return r
}

func createCubeRequest(userID string) http.Request {
	ctx := context.WithValue(context.Background(), auth.UserIDKey, userID)
	ctx = context.WithValue(ctx, auth.RoleKey, auth.RoleUser)

	body := io.NopCloser(bytes.NewReader([]byte(`
	{
		"template_name": "Test Cube",
		"template_description": "Simple cube for testing ParaCAD.",
		"template_content": "cube([10,width,10],false);",
		"template_parameters": [
			{
				"parameter_name": "width",
				"parameter_display_name": "Width of the cube",
				"parameter_type": "int",
				"parameter_default_value": "20",
				"parameter_constraints": [
					{
						"type": "min_value",
						"value": "10"
					},
					{
						"type": "max_value",
						"value": "30"
					}
				]
			}
		]
	}
	`)))

	r := http.Request{
		Method: http.MethodPost,
		Body:   body,
	}
	r = *r.WithContext(ctx)

	return r
}

func createMeasuringCupRequest(userID string) http.Request {
	ctx := context.WithValue(context.Background(), auth.UserIDKey, userID)
	ctx = context.WithValue(ctx, auth.RoleKey, auth.RoleUser)

	body := io.NopCloser(bytes.NewReader([]byte(`
	{
		"template_name": "Miarka",
		"template_description": "Kubek / miarka / naczynie do mierzenia dokładnych objętości cieczy lub proszków.",
		"template_content": "$fa=2;\n$fs=0.2;\n$fn=1000;\n\nvolume_mm3=volume*10*10*10;\n\nvoid_radius=void_diameter/2;\nvoid_height=(4*volume_mm3)/(3.1415*void_diameter*void_diameter);\n\ndifference() {\n    cylinder(void_height,\n        void_radius+wall_thickness,\n        void_radius+wall_thickness);\n    translate([0,0,wall_thickness]){\n        cylinder(void_height*1.5,void_radius,void_radius);\n    };\n}\n",
		"template_parameters": [
			{
				"parameter_name": "volume",
				"parameter_display_name": "Objętość (cm3)",
				"parameter_type": "float",
				"parameter_default_value": "4",
				"parameter_constraints": [
					{
						"type": "min_value",
						"value": "2"
					},
					{
						"type": "max_value",
						"value": "80"
					},
					{
						"type": "step",
						"value": "0.1"
					}
				]
			},
			{
				"parameter_name": "wall_thickness",
				"parameter_display_name": "Grubość ścian (mm)",
				"parameter_type": "float",
				"parameter_default_value": "1.2",
				"parameter_constraints": [
					{
						"type": "min_value",
						"value": "0.8"
					},
					{
						"type": "max_value",
						"value": "5"
					},
					{
						"type": "step",
						"value": "0.2"
					}
				]
			},
			{
				"parameter_name": "void_diameter",
				"parameter_display_name": "Średnica wewnętrzna (mm)",
				"parameter_type": "float",
				"parameter_default_value": "20",
				"parameter_constraints": [
					{
						"type": "min_value",
						"value": "10"
					},
					{
						"type": "max_value",
						"value": "80"
					},
					{
						"type": "step",
						"value": "0.5"
					}
				]
			}
		]
	}
	`)))

	r := http.Request{
		Method: http.MethodPost,
		Body:   body,
	}
	r = *r.WithContext(ctx)

	return r
}

func createBoxWithLidRequest(userID string) http.Request {
	ctx := context.WithValue(context.Background(), auth.UserIDKey, userID)
	ctx = context.WithValue(ctx, auth.RoleKey, auth.RoleUser)

	body := io.NopCloser(bytes.NewReader([]byte(`
	{
		"template_name": "Pudełko z suwanym wieczkiem",
		"template_description": "Pudełko z suwanym, demontowalnym wieczkiem. Pozwala na łatwy dostęp do bezpiecznie przechowywanej zawartości. Ustaw parametry, wygeneruj pudełko, zaznacz \"Wieczko\", wygeneruj wieczko z identycznymi parametrami. Wszystkie wartości są w mm.",
		"template_content": "if (!lid) {\n    difference(){\n        // box\n        cube([content_length+2*wall_thickness,\n        content_width+2*wall_thickness,\n        content_height+2*wall_thickness]);\n        // insides\n        translate([wall_thickness,wall_thickness,wall_thickness])\n            cube([content_length,\n            content_width,\n            content_height+2*wall_thickness]);\n        // lid\n        union(){\n            translate([wall_thickness*2/3,\n            wall_thickness/2-clearance,\n            wall_thickness+content_height])\n                cube([content_length+wall_thickness*2,\n                content_width+2*((wall_thickness/2)+clearance),\n                lid_thickness+clearance]);\n            translate([content_length+wall_thickness*0.9,\n            wall_thickness,\n            wall_thickness+content_height])\n                cube([wall_thickness*2,\n                content_width,\n                wall_thickness*2]);\n        }\n    };\n};\n\nif (lid) {\n    union() {\n        translate([wall_thickness*2/3,\n        wall_thickness/2-clearance/2,\n        wall_thickness+content_height+clearance/2]){\n            cube([content_length+wall_thickness*4/3,\n            content_width+2*(wall_thickness/2+clearance/2),\n            lid_thickness]);\n            \n        };\n        translate([content_length+wall_thickness,\n        wall_thickness+clearance/2,\n        wall_thickness+content_height+clearance/2+0.05]){\n            cube([wall_thickness-0.05,\n            content_width-clearance,\n            wall_thickness-clearance/2-0.05]);\n            \n        };\n    };\n};",
		"template_parameters": [
			{
				"parameter_name": "content_length",
				"parameter_display_name": "Długość zawartości",
				"parameter_type": "int",
				"parameter_default_value": "139",
				"parameter_constraints": [
					{
						"type": "min_value",
						"value": "15"
					},
					{
						"type": "max_value",
						"value": "200"
					}
				]
			},
			{
				"parameter_name": "content_width",
				"parameter_display_name": "Szerokość zawartości",
				"parameter_type": "int",
				"parameter_default_value": "70",
				"parameter_constraints": [
					{
						"type": "min_value",
						"value": "30"
					},
					{
						"type": "max_value",
						"value": "100"
					}
				]
			},
			{
				"parameter_name": "content_height",
				"parameter_display_name": "Wysokość zawartości",
				"parameter_type": "int",
				"parameter_default_value": "15",
				"parameter_constraints": [
					{
						"type": "min_value",
						"value": "10"
					},
					{
						"type": "max_value",
						"value": "100"
					}
				]
			},
			{
				"parameter_name": "wall_thickness",
				"parameter_display_name": "Grubość ścianek",
				"parameter_type": "float",
				"parameter_default_value": "3",
				"parameter_constraints": [
					{
						"type": "min_value",
						"value": "2"
					},
					{
						"type": "max_value",
						"value": "10"
					},
					{
						"type": "step",
						"value": "0.2"
					}
				]
			},
			{
				"parameter_name": "lid_thickness",
				"parameter_display_name": "Grubość wieczka",
				"parameter_type": "float",
				"parameter_default_value": "1.6",
				"parameter_constraints": [
					{
						"type": "min_value",
						"value": "0.4"
					},
					{
						"type": "max_value",
						"value": "5"
					},
					{
						"type": "step",
						"value": "0.4"
					}
				]
			},
			{
				"parameter_name": "clearance",
				"parameter_display_name": "Tolerancja",
				"parameter_type": "float",
				"parameter_default_value": "0.2",
				"parameter_constraints": [
					{
						"type": "min_value",
						"value": "0"
					},
					{
						"type": "max_value",
						"value": "1"
					},
					{
						"type": "step",
						"value": "0.05"
					}
				]
			},
			{
				"parameter_name": "lid",
				"parameter_display_name": "Wieczko",
				"parameter_type": "bool",
				"parameter_default_value": "false",
				"parameter_constraints": []
			}
		]
	}
	`)))

	r := http.Request{
		Method: http.MethodPost,
		Body:   body,
	}
	r = *r.WithContext(ctx)

	return r
}

func createTestTubeHolderRequest(userID string) http.Request {
	ctx := context.WithValue(context.Background(), auth.UserIDKey, userID)
	ctx = context.WithValue(ctx, auth.RoleKey, auth.RoleUser)

	body := io.NopCloser(bytes.NewReader([]byte(`
	{
		"template_name": "Uchwyt na probówki",
		"template_description": "Stojak na probówki. Przydatny podczas pracy z wieloma substancjami w laboratorium.",
		"template_content": "$fa = 5;\n$fs = 0.05;\n\nsingle_cube_size=test_tube_diameter+wall_thickness*2;\ncube_width=single_cube_size*width;\ncube_length=single_cube_size*length;\nhole_height=height-wall_thickness;\ntest_tube_radius=test_tube_diameter/2;\n\ndifference(){\n    cube([cube_width,cube_length,height]);\n    for(x = [1 : length])\n    {\n        for(y = [1 : width])\n        {\n            translate([single_cube_size/2+(y-1)*single_cube_size,\n                single_cube_size/2+(x-1)*single_cube_size,wall_thickness+0.5])\n                cylinder(hole_height,test_tube_radius,test_tube_radius);\n        };\n    };\n}",
		"template_parameters": [
			{
				"parameter_name": "test_tube_diameter",
				"parameter_display_name": "Średnica probówki (mm)",
				"parameter_type": "int",
				"parameter_default_value": "15",
				"parameter_constraints": [
					{
						"type": "min_value",
						"value": "8"
					},
					{
						"type": "max_value",
						"value": "30"
					}
				]
			},
			{
				"parameter_name": "wall_thickness",
				"parameter_display_name": "Grubość ścian (mm)",
				"parameter_type": "float",
				"parameter_default_value": "1",
				"parameter_constraints": [
					{
						"type": "min_value",
						"value": "0.2"
					},
					{
						"type": "max_value",
						"value": "5"
					},
					{
						"type": "step",
						"value": "0.2"
					}
				]
			},
			{
				"parameter_name": "height",
				"parameter_display_name": "Wysokość (mm)",
				"parameter_type": "int",
				"parameter_default_value": "20",
				"parameter_constraints": [
					{
						"type": "min_value",
						"value": "15"
					},
					{
						"type": "max_value",
						"value": "30"
					}
				]
			},
			{
				"parameter_name": "width",
				"parameter_display_name": "Szerokość (liczba probówek)",
				"parameter_type": "int",
				"parameter_default_value": "5",
				"parameter_constraints": [
					{
						"type": "min_value",
						"value": "2"
					},
					{
						"type": "max_value",
						"value": "30"
					}
				]
			},
			{
				"parameter_name": "length",
				"parameter_display_name": "Długość (liczba probówek)",
				"parameter_type": "int",
				"parameter_default_value": "2",
				"parameter_constraints": [
					{
						"type": "min_value",
						"value": "2"
					},
					{
						"type": "max_value",
						"value": "30"
					}
				]
			}
		]
	}
	`)))

	r := http.Request{
		Method: http.MethodPost,
		Body:   body,
	}
	r = *r.WithContext(ctx)

	return r
}
