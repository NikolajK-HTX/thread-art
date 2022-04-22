use std::env;
use image::io::Reader as ImageReader;
use std::io::Cursor;

fn main() {
    let args: Vec<String> = env::args().collect();
    println!("Hello, world!");
    println!("Arguments: {:?}", args);

    let img = ImageReader::open("selfie.jpg")?.decode()?;
    let img2 = ImageReader::new(Cursor::new(bytes)).with_guessed_form()?.decode()?;
   
}
