#include <iostream>
#include <stdlib.h>
#include "libbmp.h"
#include <cstring>
#include <chrono> 

struct rgb
{
    uint8_t r, g, b;
};

const int NUM_MUESTRAS = 100;

// Uso: ./invertir ruta_origen ruta_destino version
int main(int argc, char *argv[])
{
    if (argc < 4)
    {
        std::cout << "Uso: ./invertir ruta_origen ruta_destino version\nPresiona ENTER para continuar...\n";
        getchar();
        return 0;
    }

    // Lee el bmp dentro de la matriz de pixeles
    BmpImg img;
    img.read(argv[1]);
    const int width = img.get_width();
    const int height = img.get_height();
    rgb ImRGB0[height][width];

    for (int r = 0; r < height; r++)
    {
        for (int c = 0; c < width; c++)
        {
            ImRGB0[r][c].r = img.red_at(c, r);
            ImRGB0[r][c].g = img.green_at(c, r);
            ImRGB0[r][c].b = img.blue_at(c, r);
        }
    }

    rgb ImRGB[height][width];
    memcpy(ImRGB, ImRGB0, height * width * sizeof(rgb));

    //Ejecuta el algoritmo
    long version = strtol(argv[3], NULL, 10);

    freopen("out.txt", "w", stdout);

    int n = NUM_MUESTRAS;
    while (n--)
    {
        auto start = std::chrono::high_resolution_clock::now(); 
        switch (version)
        {
        case 1:
            for (int r = 0; r < height; r++)
            {
                for (int c = 0; c < width; c++)
                {
                    ImRGB0[r][c].r = 255 - ImRGB[r][c].r;
                    ImRGB0[r][c].g = 255 - ImRGB[r][c].g;
                    ImRGB0[r][c].b = 255 - ImRGB[r][c].b;
                }
            }
            break;

        default:
            break;
        }
        auto stop = std::chrono::high_resolution_clock::now(); 
    }

    //Escribe la matriz de pixeles en el nuevo bmp
    for (int r = 0; r < height; r++)
    {
        for (int c = 0; c < width; c++)
        {
            img.set_pixel(c, r, ImRGB0[r][c].r, ImRGB0[r][c].g, ImRGB0[r][c].b);
        }
    }
    img.write(argv[2]);

    return 0;
}
