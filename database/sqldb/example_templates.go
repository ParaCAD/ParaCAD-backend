package sqldb

const exampleTemplateCube = `
cube([30,width,10],false);
`

const exampleTemplateBox = `
if (!lid) {
    difference(){
        // box
        cube([content_length+2*wall_thickness,
        content_width+2*wall_thickness,
        content_height+2*wall_thickness]);
        // insides
        translate([wall_thickness,wall_thickness,wall_thickness])
            cube([content_length,
            content_width,
            content_height+2*wall_thickness]);
        // lid
        union(){
            translate([wall_thickness*2/3,
            wall_thickness/2-clearance,
            wall_thickness+content_height])
                cube([content_length+wall_thickness*2,
                content_width+2*((wall_thickness/2)+clearance),
                lid_thickness+clearance]);
            translate([content_length+wall_thickness*0.9,
            wall_thickness,
            wall_thickness+content_height])
                cube([wall_thickness*2,
                content_width,
                wall_thickness*2]);
        }
    };
};

if (lid) {
    union() {
        translate([wall_thickness*2/3,
        wall_thickness/2-clearance/2,
        wall_thickness+content_height+clearance/2]){
            cube([content_length+wall_thickness*4/3,
            content_width+2*(wall_thickness/2+clearance/2),
            lid_thickness]);
            
        };
        translate([content_length+wall_thickness,
        wall_thickness+clearance/2,
        wall_thickness+content_height+clearance/2+0.05]){
            cube([wall_thickness-0.05,
            content_width-clearance,
            wall_thickness-clearance/2-0.05]);
            
        };
    };
};
`
