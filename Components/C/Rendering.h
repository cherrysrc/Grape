#include <stdio.h>
#include <string.h>
#include <stdlib.h>

typedef struct
{
    int width;
    int height;
    unsigned char* buffer;
} Rendering;

Rendering* createRendering(int width, int height)
{
    Rendering* rendering = (Rendering*)calloc(1, sizeof(Rendering));

    rendering->buffer = (unsigned char*)calloc(3 * width * height, sizeof(unsigned char));
    rendering->width  = width;
    rendering->height = height;

    return(rendering);
}

void setPixel(Rendering* rendering, int x, int y, unsigned char r, unsigned char g, unsigned char b)
{
    rendering->buffer[(x + y * rendering->width) * 3 + 0] = r;
    rendering->buffer[(x + y * rendering->width) * 3 + 1] = g;
    rendering->buffer[(x + y * rendering->width) * 3 + 2] = b;
}

void saveRendering(Rendering* rendering, char* filename)
{
    FILE* fp;

    fp = fopen(filename, "w");

    if (fp == NULL)
    {
        printf("Can't open file!");
        return;
    }

    fprintf(fp, "P6\n%d %d\n255\n", rendering->width, rendering->height);
    fwrite(rendering->buffer, 3 * rendering->width * rendering->height, 1, fp);

    fclose(fp);
}

void writeRendering(Rendering* rendering)
{
    printf("P6\n%d %d\n255\n", rendering->width, rendering->height);
    fwrite(rendering->buffer, sizeof(unsigned char), 3 * rendering->width * rendering->height, stdout);
    fflush(stdout);
}

void freeRendering(Rendering* rendering)
{
    free(rendering->buffer);
    free(rendering);
}
