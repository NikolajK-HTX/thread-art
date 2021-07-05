using System;
using System.Drawing;

namespace threadArtApplication
{
    class Program
    {
        const string INPUT_IMAGE_PATH = @"C:\Users\nikol\Documents\GitHub\thread-art\myselfie.jpg";

        const string OUTPUT_IMAGE_PATH = @"C:\Users\nikol\Documents\GitHub\thread-art\";
        const int OUTPUT_IMAGE_SIZE = 400;

        static void Main(string[] args)
        {
            Console.WriteLine("Hello World!");

            Bitmap inputImage = new Bitmap(INPUT_IMAGE_PATH);

            int x, y;
            // Loop through the images pixels to reset color.
            for(x=0; x<inputImage.Width; x++)
            {
                for(y=0; y<inputImage.Height; y++)
                {
                    Color pixelColor = inputImage.GetPixel(x, y);
                    Color newColor = Color.FromArgb(pixelColor.R, 0, 0);
                    inputImage.SetPixel(x, y, newColor);
                }
            }
            // save the manipulated image
            inputImage.Save(@"C:\Users\nikol\Documents\GitHub\thread-art\bob.jpg");

            Bitmap outputImage = new Bitmap(OUTPUT_IMAGE_SIZE, OUTPUT_IMAGE_SIZE);
            Graphics g = Graphics.FromImage(outputImage);
        }
    }
}
