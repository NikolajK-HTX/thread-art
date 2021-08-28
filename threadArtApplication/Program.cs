using System;
using SixLabors.ImageSharp;
using SixLabors.ImageSharp.PixelFormats;
using System.IO;
using System.Collections.Generic;
using System.Runtime.InteropServices;

//

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
        int centerX;
        int centerY;
        int radius;
        public List<Point> points;
        public Dictionary<string, int[,]> allLines = new Dictionary<string, int[,]>();
        int index;

        public Circle(int center_x, int center_y, int _radius, int num_points)
        {
            centerX = center_x;
            centerY = center_y;
            radius = _radius;
            points = new List<Point>();
            index = 0;
            for (int i = 0; i < num_points; i++)
            {
                int x = (int)Math.Min((Math.Cos(Math.PI * 2 * i / num_points) * radius + centerX), radius * 2 - 1);
                int y = (int)Math.Min((Math.Sin(Math.PI * 2 * i / num_points) * radius + centerY), radius * 2 - 1);
                points.Add(new Point(x, y, index));
                index++;
            }
        }

        public int[] getXY(int _index)
        {
            return (new int[2] { points[_index].x, points[_index].y });
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
            }
            else
            {
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

            int[,] pixels = new int[dx + 1, 2];

            for (int x = 0; x < dx + 1; x++)
            {
                pixels[x, 0] = x0 + x * xx + y * yx;
                pixels[x, 1] = y0 + x * xy + y * yy;
                if (D > 0)
                {
                    y++;
                    D -= dx;
                }
                D += dy;
            }

            return pixels;
        }

        static int lineWeight(byte[] pixels, int[,] line, BitmapData bitmapData, int widthInBytes)
        {
            int sum = line.GetLength(0) * 255;
            for (int subArray = 0; subArray < line.GetLength(0); subArray++)
            {
                int x = line[subArray, 0];
                int y = line[subArray, 1];
                sum -= GetBrightnessFromPixel(x, y, pixels, bitmapData, widthInBytes);
            }
            return (sum / line.GetLength(0));
        }

        static int GetBrightnessFromPixel(int x, int y, byte[] pixels, BitmapData bitmapData, int bytesPerPixel)
        {
            int currentLine = y * bitmapData.Stride;
            int R = pixels[currentLine + x * bytesPerPixel];
            int G = pixels[currentLine + x * bytesPerPixel + 1];
            int B = pixels[currentLine + x * bytesPerPixel + 2];

            float r = (float)R / 255.0f;
            float g = (float)G / 255.0f;
            float b = (float)B / 255.0f;

            float max, min;

            max = r; min = r;

            if (g > max) max = g;
            if (b > max) max = b;

            if (g < min) min = g;
            if (b < min) min = b;

            return (int)(((max + min) / 2) * 255);
        }

        static void SetPixelBrightness(byte value, int x, int y, byte[] pixels, BitmapData bitmapData, int bytesPerPixel)
        {
            int currentLine = y * bitmapData.Stride;
            pixels[currentLine + x * bytesPerPixel] = value;
            pixels[currentLine + x * bytesPerPixel + 1] = value;
            pixels[currentLine + x * bytesPerPixel + 2] = value;
        }

        static void changeBrightness(byte[] pixels, int[,] line, BitmapData bitmapData, int bytesPerPixel)
        {
            for (int subArrayCounter = 0; subArrayCounter < line.GetLength(0); subArrayCounter++)
            {
                int x = line[subArrayCounter, 0];
                int y = line[subArrayCounter, 1];

                int value = GetBrightnessFromPixel(x, y, pixels, bitmapData, bytesPerPixel);
                value += BRIGHTNESS_INCREASE_VALUE;
                value = value > 255 ? 255 : value;

                SetPixelBrightness((byte)value, x, y, pixels, bitmapData, bytesPerPixel);
            }
        }

        static string pair(int a, int b)
        {
            return (a < b ? a + "-" + b : b + "-" + a);
        }

        static void linesList(int steps, Bitmap image, Circle circle, List<string> usedPoints, List<int[]> pointsList, int minimumDifference)
        {
            // locks image in memory
            BitmapData bitmapData = image.LockBits(new Rectangle(0, 0, image.Width, image.Height), ImageLockMode.ReadWrite, image.PixelFormat);

            int bytesPerPixel = Bitmap.GetPixelFormatSize(image.PixelFormat) / 8;
            int byteCount = bitmapData.Stride * image.Height;
            byte[] pixels = new byte[byteCount];
            IntPtr ptrFirstPixel = bitmapData.Scan0;
            Marshal.Copy(ptrFirstPixel, pixels, 0, pixels.Length);
            int heightInPixels = bitmapData.Height;

            Point startPoint = circle.points[0];
            for (int loops = 0; loops < steps; loops++)
            {
                int maxWeight = 0;
                Point nextPoint = new Point(0, 0, 0);
                int[,] maxLine = new int[,] { { 0, 0 }, { 0, 0 } };
                foreach (Point point in circle.points)
                {
                    int difference = Math.Abs(point.index - startPoint.index);
                    if (difference < minimumDifference || difference > (circle.points.Count - minimumDifference))
                    {
                        continue;
                    }
                    int weight = lineWeight(pixels, circle.allLines[pair(startPoint.index, point.index)], bitmapData, bytesPerPixel);
                    if (weight > maxWeight && point != startPoint && !usedPoints.Contains(pair(startPoint.index, point.index)))
                    {
                        maxWeight = weight;
                        nextPoint = point;
                        maxLine = circle.allLines[pair(startPoint.index, point.index)];
                    }
                }
                usedPoints.Add(pair(startPoint.index, nextPoint.index));
                pointsList.Add(new int[2] { startPoint.index, nextPoint.index });
                changeBrightness(pixels, maxLine, bitmapData, bytesPerPixel);
                startPoint = nextPoint;
            }

            Marshal.Copy(pixels, 0, ptrFirstPixel, pixels.Length);
            image.UnlockBits(bitmapData);
        }

        static Bitmap draw(List<int[]> pointsList, Circle circle, int size)
        {
            Bitmap image = new Bitmap(size, size);
            Graphics myGraphics = Graphics.FromImage(image);
            myGraphics.Clear(Color.White);
            Pen blackPen = new Pen(Color.FromArgb(100, 0, 0, 0), 2);
            for (int i = 0; i < pointsList.Count; i++)
            {
                int[] temporaryPoint = pointsList[i];
                int[] firstPointArray = circle.getXY(temporaryPoint[0]);
                int[] secondPointArray = circle.getXY(temporaryPoint[1]);
                System.Drawing.Point firstPoint = new System.Drawing.Point(firstPointArray[0], firstPointArray[1]);
                System.Drawing.Point secondPoint = new System.Drawing.Point(secondPointArray[0], secondPointArray[1]);
                myGraphics.DrawLine(blackPen, firstPoint, secondPoint);

            }
            blackPen.Dispose();
            return (image);
        }

        static void Main(string[] args)
        {
            DateTime firstTime = DateTime.Now;
            // "Constants" - really settings - meant to be changed with arguments
            string CURRENT_DIRECTORY = Directory.GetCurrentDirectory();
            string PARENT_DIRECTORY = Directory.GetParent(CURRENT_DIRECTORY).FullName;
            string INPUT_IMAGE_PATH = Path.Combine(PARENT_DIRECTORY, "selfie.jpg");
            // above is not really needed since 
            // "new Bitmap(string)" can accept a relative path

            // The reason for the string array is because chosen settings should
            // be shown in the name of the file.
            string OUTPUT_IMAGE_FILENAME = "";
            string OUTPUT_IMAGE_PATH = "";
            string IMAGE_ID = "AAAAA";
            int OUTPUT_IMAGE_SIZE = 1024;

            int NUMBER_OF_THREADS = 2000;
            int NUMBER_OF_PINS = 200;
            int MINIMUM_DIFFERENCE = 20;

            string helpString = "./threadArtApplication -i <input_image> -n <number_of_pins> -s <outputimage_size> -t <number_of_threads> -m <minimum_difference> -o <output-image-path> -p <image_id>";

            // Change settings based on arguments from terminal
            for (int i = 0; i < args.Length; i += 2)
            {
                string arg = args[i];
                if (arg == "-h" || arg == "--help")
                {
                    Console.WriteLine(helpString);
                    Environment.Exit(0);
                }
                else if (arg == "-i" || arg == "--input-image")
                {
                    INPUT_IMAGE_PATH = args[i + 1];
                }
                else if (arg == "-t" || arg == "--number-of-threads")
                {
                    NUMBER_OF_THREADS = int.Parse(args[i + 1]);
                }
                else if (arg == "-n" || arg == "--number-of-pins")
                {
                    NUMBER_OF_PINS = int.Parse(args[i + 1]);
                }
                else if (arg == "-o" || arg == "--output-image")
                {
                    OUTPUT_IMAGE_FILENAME = args[i + 1];
                }
                else if (arg == "-s" || arg == "--output-image-size")
                {
                    OUTPUT_IMAGE_SIZE = int.Parse(args[i + 1]);
                }
                else if (arg == "-m" || arg == "--minimum-difference")
                {
                    MINIMUM_DIFFERENCE = int.Parse(args[i + 1]);
                }
                else if (arg == "-d" || arg == "--output-image-path")
                {
                    OUTPUT_IMAGE_PATH = args[i + 1];
                }
                else if (arg == "-p" || arg == "--image-id")
                {
                    IMAGE_ID = args[i + 1];
                }
            }

            // If the user has not decided where to place the output image
            // then the program sets the output itself
            string ID_FILE_NAME = String.Join('-', new String[] { IMAGE_ID, NUMBER_OF_THREADS.ToString(), NUMBER_OF_PINS.ToString(), MINIMUM_DIFFERENCE.ToString() });

            if (OUTPUT_IMAGE_PATH == "")
            {
                OUTPUT_IMAGE_FILENAME = ID_FILE_NAME + ".png";
                OUTPUT_IMAGE_PATH = Path.Combine(PARENT_DIRECTORY, OUTPUT_IMAGE_FILENAME);
            }

            // Write used settings in console
            Console.WriteLine("Hello, you are running threadArtApplication!");
            Console.WriteLine();
            Console.WriteLine("The following settings will be used!");
            Console.WriteLine("INPUT_IMAGE_PATH: " + INPUT_IMAGE_PATH);
            Console.WriteLine("NUMBER_OF_PINS: " + NUMBER_OF_PINS);
            Console.WriteLine("NUMBER_OF_THREADS: " + NUMBER_OF_THREADS);
            Console.WriteLine("OUTPUT_IMAGE_PATH: " + OUTPUT_IMAGE_PATH);
            Console.WriteLine("OUTPUT_IMAGE_SIZE: " + OUTPUT_IMAGE_SIZE);
            Console.WriteLine();

            DateTime secondTime = DateTime.Now;
            TimeSpan ts = secondTime - firstTime;
            double msTimespan = ts.TotalMilliseconds;
            Console.WriteLine(String.Format("Iniatialisering af indstillinger tog: {0} millisekunder", msTimespan));
            firstTime = DateTime.Now;

            // Load image
            if (!File.Exists(INPUT_IMAGE_PATH))
            {
                Console.WriteLine("Input image was not found :(");
                Console.WriteLine(helpString);
                Environment.Exit(2);
            }

            Image<Rgb24> inputImage = Image.Load<Rgb24>(INPUT_IMAGE_PATH);

            int imageWidth = inputImage.Width;
            int imageHeight = inputImage.Height;
            Circle imageCircle = new Circle(imageWidth / 2, imageHeight / 2, imageWidth / 2, NUMBER_OF_PINS);

            secondTime = DateTime.Now;
            ts = secondTime - firstTime;
            msTimespan = ts.TotalMilliseconds;
            Console.WriteLine(String.Format("Cirkel iniatialisering tog: {0} millisekunder", msTimespan));
            firstTime = DateTime.Now;

            // Create a dictionary that contains all possible lines
            for (int i = 0; i < NUMBER_OF_PINS; i++)
            {
                for (int j = i + 1; j < NUMBER_OF_PINS; j++)
                {
                    int[] fP = imageCircle.getXY(i);
                    int[] sP = imageCircle.getXY(j);
                    imageCircle.allLines.Add(pair(i, j), RasterLine(fP[0], fP[1], sP[0], sP[1]));
                }
            }

            secondTime = DateTime.Now;
            ts = secondTime - firstTime;
            msTimespan = ts.TotalMilliseconds;
            Console.WriteLine(String.Format("Dict med alle mulige linjer: {0} millisekunder", msTimespan));
            firstTime = DateTime.Now;


            List<int[]> pointsList = new List<int[]>();
            List<string> usedPoints = new List<string>();

            linesList(NUMBER_OF_THREADS, inputImage, imageCircle, usedPoints, pointsList, MINIMUM_DIFFERENCE);
            inputImage.Dispose();

            secondTime = DateTime.Now;
            ts = secondTime - firstTime;
            msTimespan = ts.TotalMilliseconds;
            Console.WriteLine(String.Format("Main algoritme tog: {0} millisekunder", msTimespan));
            firstTime = DateTime.Now;


            Circle outputCircle = new Circle(OUTPUT_IMAGE_SIZE / 2, OUTPUT_IMAGE_SIZE / 2, OUTPUT_IMAGE_SIZE / 2, NUMBER_OF_PINS);

            Bitmap outputImage = draw(pointsList, outputCircle, OUTPUT_IMAGE_SIZE);
            outputImage.Save(OUTPUT_IMAGE_PATH, System.Drawing.Imaging.ImageFormat.Png);

            secondTime = DateTime.Now;
            ts = secondTime - firstTime;
            msTimespan = ts.TotalMilliseconds;
            Console.WriteLine(String.Format("Tegn billede: {0} millisekunder", msTimespan));
            firstTime = DateTime.Now;

            // save pointslist to txt file
            string[] lines = new string[pointsList.Count+1];
            for (int i = 0; i < pointsList.Count; i++)
            {
                lines[i] = pointsList[i][0].ToString()+",";
                if(i == pointsList.Count-1)
                {
                    lines[i+1] = pointsList[i][1].ToString();
                }
            }
            File.WriteAllLines($"{ID_FILE_NAME}.txt", lines);

            secondTime = DateTime.Now;
            ts = secondTime - firstTime;
            msTimespan = ts.TotalMilliseconds;
            Console.WriteLine(String.Format("At skrive til txt-fil tog: {0} millisekunder", msTimespan));
        }
    }
}
