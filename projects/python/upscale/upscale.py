from super_image import EdsrModel, ImageLoader
from PIL import Image
from pathlib import Path
import os
import cv2

out_folder = "upscaled"
folders = ["input"]
model = EdsrModel.from_pretrained('eugenesiow/edsr-base', scale=2)

num_upscale = 2

Path(f'{out_folder}').mkdir(exist_ok=True)


def read_file(path):
    image = Image.open(path)
    return ImageLoader.load_image(image)


def upscale(input):
    preds = model(input)
    return ImageLoader._process_image_to_save(preds)


def out_path(file_name, target_folder):
    rest = os.path.splitext(os.path.basename(file_name))[0]
    rest = 'u-' + rest
    return f'./{out_folder}/{target_folder}/{rest}.png'


for folder in folders:
    files: list[str] = os.listdir(folder)

    Path(f'{out_folder}/{folder}').mkdir(exist_ok=True)

    for file in files:

        path = f'{folder}/{file}'
        print(f'Reading file: {path}')

        for i in range(0, num_upscale):
            print(f'Processing nth upscale ({i})')

            raw = read_file(path)
            result_1 = upscale(raw)
            path_output = out_path(path, folder)

            print(f'Saving file ({i}): {path_output}')
            cv2.imwrite(Path(path_output).__str__(), result_1)

            path = path_output
