#ifndef _OPENCV3_IMGPROC_H_
#define _OPENCV3_IMGPROC_H_

#include <stdbool.h>

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
extern "C" {
#endif

#include "core.h"

void CvtColor(Mat src, Mat dst, int code);
void BilateralFilter(Mat src, Mat dst, int d, double sc, double ss);
void Blur(Mat src, Mat dst, Size ps);
void Dilate(Mat src, Mat dst, Mat kernel);
void Erode(Mat src, Mat dst, Mat kernel);
void GaussianBlur(Mat src, Mat dst, Size ps, double sX, double sY, int bt);
Mat GetStructuringElement(int shape, Size ksize);
void MorphologyEx(Mat src, Mat dst, int op, Mat kernel);
void MedianBlur(Mat src, Mat dst, int ksize);
void Canny(Mat src, Mat edges, double t1, double t2);
void HoughCircles(Mat src, Mat circles, int method, double dp, double minDist);
void HoughLines(Mat src, Mat lines, double rho, double theta, int threshold);
void HoughLinesP(Mat src, Mat lines, double rho, double theta, int threshold);

void ArrowedLine(Mat img, Point pt1, Point pt2, Scalar color, int thickness);
void Circle(Mat img, Point center, int radius, Scalar color, int thickness);
void Line(Mat img, Point pt1, Point pt2, Scalar color, int thickness);
void Rectangle(Mat img, Rect rect, Scalar color, int thickness);
struct Size GetTextSize(const char* text, int fontFace, double fontScale, int thickness);
void PutText(Mat img, const char* text, Point org, int fontFace, double fontScale,
             Scalar color, int thickness);
void Resize(Mat src, Mat dst, Size sz, double fx, double fy, int interp);


#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_IMGPROC_H_
