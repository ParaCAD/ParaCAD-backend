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
	r := createUserRequest("test", "test@test.com", "1234")
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

	w = httptest.NewRecorder()
	r = createMeasuringCupRequest(userID.UUID.String())
	c.HandleCreateTemplate(w, &r, nil)

	w = httptest.NewRecorder()
	r = createBoxWithLidRequest(userID.UUID.String())
	c.HandleCreateTemplate(w, &r, nil)
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
		"template_name": "Measuring cup",
		"template_description": "A cup / scoop for measuring exact amounts of liquids and powders.",
		"template_content": "$fa=2;\n$fs=0.2;\n$fn=1000;\n\nvolume_mm3=volume*10*10*10;\n\nvoid_radius=void_diameter/2;\nvoid_height=(4*volume_mm3)/(3.1415*void_diameter*void_diameter);\n\ndifference() {\n    cylinder(void_height,\n        void_radius+wall_thickness,\n        void_radius+wall_thickness);\n    translate([0,0,wall_thickness]){\n        cylinder(void_height*1.5,void_radius,void_radius);\n    };\n}\n",
		"template_parameters": [
			{
				"parameter_name": "volume",
				"parameter_display_name": "Volume (cm3)",
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
				"parameter_display_name": "Wall thickness (mm)",
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
				"parameter_display_name": "Inner diameter (mm)",
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
		"template_name": "Box with sliding lid",
		"template_description": "Simple box with sliding lid. Lid is not attached to the box, allowing for easy access to the contents. Set parameters, generate box, check 'Generate lid', generate lid. All dimensions are in mm.",
		"template_content": "if (!lid) {\n    difference(){\n        // box\n        cube([content_length+2*wall_thickness,\n        content_width+2*wall_thickness,\n        content_height+2*wall_thickness]);\n        // insides\n        translate([wall_thickness,wall_thickness,wall_thickness])\n            cube([content_length,\n            content_width,\n            content_height+2*wall_thickness]);\n        // lid\n        union(){\n            translate([wall_thickness*2/3,\n            wall_thickness/2-clearance,\n            wall_thickness+content_height])\n                cube([content_length+wall_thickness*2,\n                content_width+2*((wall_thickness/2)+clearance),\n                lid_thickness+clearance]);\n            translate([content_length+wall_thickness*0.9,\n            wall_thickness,\n            wall_thickness+content_height])\n                cube([wall_thickness*2,\n                content_width,\n                wall_thickness*2]);\n        }\n    };\n};\n\nif (lid) {\n    union() {\n        translate([wall_thickness*2/3,\n        wall_thickness/2-clearance/2,\n        wall_thickness+content_height+clearance/2]){\n            cube([content_length+wall_thickness*4/3,\n            content_width+2*(wall_thickness/2+clearance/2),\n            lid_thickness]);\n            \n        };\n        translate([content_length+wall_thickness,\n        wall_thickness+clearance/2,\n        wall_thickness+content_height+clearance/2+0.05]){\n            cube([wall_thickness-0.05,\n            content_width-clearance,\n            wall_thickness-clearance/2-0.05]);\n            \n        };\n    };\n};",
		"template_parameters": [
			{
				"parameter_name": "content_length",
				"parameter_display_name": "Content length",
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
				"parameter_display_name": "Content width",
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
				"parameter_display_name": "Content height",
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
				"parameter_display_name": "Wall thickness",
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
				"parameter_display_name": "Lid thickness",
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
				"parameter_display_name": "Clearance",
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
				"parameter_display_name": "Generate lid",
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
