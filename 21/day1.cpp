#include <fstream>
#include <iostream>

int main() {
  int num, curr = 0, prev = 0;
  int index = 0, counter = 0;

  std::ifstream infile("./inputs/1.txt"); 

  while(infile >> num){
    if(index++ < 3) {
      curr += num;
    } else {
      if(prev != 0) counter += curr > prev;

      index = 0;
      prev  = curr;
    }


    // part 1
    //if(prev != 0) counter +=  curr > prev ;
    //prev = curr;
  }

  std::cout<<counter;
}
