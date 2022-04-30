use image::{io::Reader, GenericImageView};
use line_drawing::Bresenham;
use std::process;
use std::f64::consts::PI;
use std::env;

struct Point {
    x: u32,
    y: u32
}

struct Pair {
    a: u32,
    b: u32,
    points: Vec<Point>
}

struct Circle {
    x_center: u32,
    y_center: u32,
    points: Vec<Point>
}

fn main() {
    let args: Vec<String> = env::args().collect();

    let image = Reader::open("selfie.jpg").unwrap()
        .decode().unwrap();

    let width = image.width();
    let height = image.height();

    if width != height {
        process::exit(1);
    }

    let mut num_points: u32 = args[1].parse().unwrap();

    let mut sum: u32 = 0;
    for x in 0..width {
        for y in 0..height {
            let pixel = image.get_pixel(x, y);
            sum += pixel[0] as u32;
        }
    }
    println!("{:?}", sum);

    let mut circle = Circle {
        x_center: width/2,
        y_center: height/2,
        points: Vec::new(),
    };

    let step_angle = PI*2.0/num_points as f64;
    for i in 0..num_points {
        let angle = step_angle*i as f64;
        let x = (angle.cos()*width as f64/2.0).floor() as u32;
        let y = (angle.sin()*width as f64/2.0).floor() as u32;
        let point = Point {x, y};
        circle.points.push(point);
    }

    let mut points_pair: Vec<(u32, u32)> = Vec::new();
    let mut counter = 0;
    for i in 0..num_points {
        for n in i+1..num_points {
            let start_point = &circle.points[i as usize];
            let end_point = &circle.points[n as usize];
            let start_point = (start_point.x as i32, start_point.y as i32);
            let end_point = (end_point.x as i32, end_point.y as i32);
            counter += 1;
            for (x, y) in Bresenham::new(start_point, end_point) {
                points_pair.push((x as u32, y as u32));
            }
        }
    }

    // println!("{:?}", points_pair.len());
    println!("{:?}", counter);


    // let x = 12;
    // let y = 24;

    // let pixel = image.get_pixel(x, y);

    // println!("{:?}", pixel);

    // let width = image.width();
    // let _height = image.height();

    // let bytes = image.into_bytes();
    // //println!("{:?}", &bytes[1..10]);

    // let position: usize = (width*3*y+width*x).try_into().unwrap();
    // println!("{:?}", &bytes[position..position+3]);
}