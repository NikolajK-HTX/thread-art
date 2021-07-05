using System;
using System.Drawing;
using System.IO;

namespace threadArtApplication
{
    class Program
    {
        static void Main(string[] args)
        {
            string CURRENT_DIRECTORY = Directory.GetCurrentDirectory();
            string PARENT_DIRECTORY = Directory.GetParent(CURRENT_DIRECTORY).FullName;
            string INPUT_IMAGE_PATH = Path.Combine(PARENT_DIRECTORY, "myselfie.jpg");

            string OUTPUT_IMAGE_PATH = Path.Combine(PARENT_DIRECTORY, "outputimage.jpg");
            const int OUTPUT_IMAGE_SIZE = 400;

            Console.WriteLine("Hello World!");
            Console.WriteLine("OUTPUT_IMAGE_PATH: "+OUTPUT_IMAGE_PATH);
            Console.WriteLine("INPUT_IMAGE_PATH: "+INPUT_IMAGE_PATH);

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
            inputImage.Save(OUTPUT_IMAGE_PATH);

            Bitmap outputImage = new Bitmap(OUTPUT_IMAGE_SIZE, OUTPUT_IMAGE_SIZE);
            Graphics g = Graphics.FromImage(outputImage);
        }
    }
}
