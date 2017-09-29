
#include <iostream>
using namespace std;

class CRectangle {
    int x, y;
  	void setValues(int, int);
  public:
  	CRectangle(int, int);
  	int area();
};

CRectangle::CRectangle(int a, int b) {
	setValues(a, b);
}

void CRectangle::setValues(int a, int b) {
	x = a;
	y = b;
}

int CRectangle::area() {
	return x * y;
}


int main() {

  int a,b,c;
  string mystr;

  a = 2;
  b = 7;

  c = (a > b) ? a : b;

  cout << "Is your age less than " << c << "?\n";
  getline (cin, mystr);
  cout << "That's OK!\n";

  CRectangle rect(3, 6);

//  rect.setValues(2, 5);

  cout << "Area: " << rect.area() << "\n";

  return 0;
 }
