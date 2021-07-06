using System;
using System.Drawing;
using System.IO;
using System.Collections.Generic;

namespace threadArtApplication
{
    class Point
    {
        int x;
        int y;
        int index;
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
    }

    class Program
    {
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

        static void Main(string[] args)
        {
            string CURRENT_DIRECTORY = Directory.GetCurrentDirectory();
            string PARENT_DIRECTORY = Directory.GetParent(CURRENT_DIRECTORY).FullName;
            string INPUT_IMAGE_PATH = Path.Combine(PARENT_DIRECTORY, "myselfie.jpg");

            string OUTPUT_IMAGE_FILENAME = String.Join('-', new String[] {"outputimage"})+".jpg";
            string OUTPUT_IMAGE_PATH = Path.Combine(PARENT_DIRECTORY, OUTPUT_IMAGE_FILENAME);
            const int OUTPUT_IMAGE_SIZE = 400;

            Console.WriteLine("Hello World!");
            Console.WriteLine("OUTPUT_IMAGE_PATH: "+OUTPUT_IMAGE_PATH);
            Console.WriteLine("INPUT_IMAGE_PATH: "+INPUT_IMAGE_PATH);

            Bitmap inputImage = new Bitmap(INPUT_IMAGE_PATH);

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

            Bitmap outputImage = new Bitmap(OUTPUT_IMAGE_SIZE, OUTPUT_IMAGE_SIZE);
            Graphics g = Graphics.FromImage(outputImage);
        }
    }
}
