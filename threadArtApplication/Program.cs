using System;
using System.Drawing;
using System.IO;
using System.Collections.Generic;

namespace threadArtApplication
{
    class Point
    {
        public int x;
        public int y;
        public int index;
        public Point(int _x, int _y, int _index)
        {
            x = _x;
            y = _y;
            index = _index;
        }
    }

    class Circle
    {
        float angle;
        int centerX;
        int centerY;
        int radius;
        float step;
        List<Point> points;
        int index;

        public Circle(int center_x, int center_y, int _radius, int num_points)
        {
            centerX = center_x;
            centerY = center_y;
            radius = _radius;
            step = 360 / num_points;
            points = new List<Point>();
            index = 0;
            for (angle = 0; angle < 360; angle += step)
            {
                int x = (int)Math.Cos(angle)*radius+centerX;
                int y = (int)Math.Sin(angle)*radius+centerY;
                points.Add(new Point(x, y, index));
                index++;
            }
        }

        int[] getXY(int _index) 
        {
            return (new int[2] {points[_index].x, points[_index].y});
        }
    }

    class Program
    {
        const int BRIGHTNESS_INCREASE_VALUE = 50;

        static int[,] RasterLine(int x0, int y0, int x1, int y1) 
        {
            int dx = x1 - x0;
            int dy = y1 - y0;

            int xsign = dx > 0 ? 1 : -1;
            int ysign = dy > 0 ? 1 : -1;

            dx = Math.Abs(dx);
            dy = Math.Abs(dy);

            int xx, xy, yx, yy;
            if (dx > dy) 
            {
                xx = xsign;
                xy = 0;
                yx = 0;
                yy = ysign;
            } else {
                int temp = dx;
                dx = dy;
                dy = temp;
                xx = 0;
                xy = ysign;
                yx = xsign;
                yy = 0;
            }

            int D = 2 * dy - dx;
            int y = 0;

            int[,] pixels = new int[dx+1, 2];

            for (int x = 0; x < dx+1; x++)
            {
                pixels[x, 0] = x0 +x * xx + y * yx;
                pixels[x, 1] = y0 +x * xy + y * yy;
                if (D>0) {
                    y++;
                    D -= dx;
                }
                D += dy;
            }

            return pixels;
        }

        static int lineWeight(Bitmap image, int[,] line) 
        {
            int sum = line.GetLength(0) * 255;
            for (int subArray = 0; subArray < line.GetLength(0); subArray++)
            {
                int x = line[subArray, 0];
                int y = line[subArray, 1];
                sum -= (int)(image.GetPixel(x, y).GetBrightness()*255);
            }
            return (sum);
        }

        static void changeBrightness(ref Bitmap image, int[,] line) 
        {
            for (int subArray = 0; subArray < line.GetLength(0); subArray++)
            {
                int x = line[subArray, 0];
                int y = line[subArray, 1];

                int value = (int) (image.GetPixel(x, y).GetBrightness()*255);
                value += BRIGHTNESS_INCREASE_VALUE;
                value = value > 255 ? 255 : value;

                image.SetPixel(x, y, Color.FromArgb(value, value, value));
            }
        }

        static string pair (int a, int b)
        {
            return(a<b ? String.Join('-', new int[2] {a, b}) : String.Join('-', new int[2] {b, a}));
        }

        static void linesList() 
        {

        }

        static void Main(string[] args)
        {
            // "Constants" - really settings - meant to be changed with arguments
            string CURRENT_DIRECTORY = Directory.GetCurrentDirectory();
            string PARENT_DIRECTORY = Directory.GetParent(CURRENT_DIRECTORY).FullName;
            string INPUT_IMAGE_PATH = Path.Combine(PARENT_DIRECTORY, "myselfie.jpg"); 
            // above is not really needed since 
            // "new Bitmap(string)" can accept a relative path

            string OUTPUT_IMAGE_FILENAME = String.Join('-', new String[] {"outputimage"})+".jpg";
            string OUTPUT_IMAGE_PATH = Path.Combine(PARENT_DIRECTORY, OUTPUT_IMAGE_FILENAME);
            int OUTPUT_IMAGE_SIZE = 400;

            int NUMBER_OF_THREADS = 2000;
            int NUMBER_OF_PINS = 200;

            // Change settings based on arguments from terminal
            for (int i = 0; i < args.Length; i += 2)
            {
                string arg = args[i];
                if (arg == "-i" || arg == "--input-image") {
                    INPUT_IMAGE_PATH = args[i+1];
                } else if (arg == "-t" || arg == "--number-of-threads") {
                    NUMBER_OF_THREADS = int.Parse(args[i+1]);
                } else if (arg == "-n" || arg == "--number-of-pins") {
                    NUMBER_OF_PINS = int.Parse(args[i+1]);
                } else if (arg == "-o" || arg == "--output-image") {
                    OUTPUT_IMAGE_FILENAME = args[i+1];
                }
            }

            // Write used settings in console
            Console.WriteLine("Hello, you are running threadArtApplication!");
            Console.WriteLine();
            Console.WriteLine("The following settings will be used!");
            Console.WriteLine("OUTPUT_IMAGE_PATH: "+OUTPUT_IMAGE_PATH);
            Console.WriteLine("INPUT_IMAGE_PATH: "+INPUT_IMAGE_PATH);
            Console.WriteLine("NUMBER_OF_PINS: "+NUMBER_OF_PINS);
            Console.WriteLine("NUMBER_OF_THREADS: "+NUMBER_OF_THREADS);
            Console.WriteLine();

            // Load image
            if (!File.Exists(INPUT_IMAGE_PATH)) 
            {
                Console.WriteLine("Input image was not found :(");
                Environment.Exit(2);
            }
            /* Bitmap inputImage = new Bitmap(INPUT_IMAGE_PATH);   

            int imageWidth = inputImage.Width;
            int imageHeight = inputImage.Height;
            Circle imageCircle = new Circle(imageWidth/2, imageHeight/2, imageWidth/2, 200);
            Console.WriteLine(imageCircle);

            int x, y;
            // Loop through the images pixels and only use red color value.
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

            // Create a new image from Graphics
            Bitmap outputImage = new Bitmap(OUTPUT_IMAGE_SIZE, OUTPUT_IMAGE_SIZE);
            Graphics g = Graphics.FromImage(outputImage); */
        }
    }
}
