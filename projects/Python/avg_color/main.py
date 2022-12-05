import cv2, numpy as np
from sklearn.cluster import KMeans

image_data = [
    ("Albedo", "albedo_gacha_splash.png"),
    ("Aloy", "aloy_gacha_splash.png"),
    ("Amber", "amber_gacha_splash.png"),
    ("Arataki Itto", "itto_gacha_splash.png"),
    ("Barbara", "barbara_gacha_splash.png"),
    ("Beidou", "beidou_gacha_splash.png"),
    ("Bennet", "bennett_gacha_splash.png"),
    ("Chongyun", "chongyun_gacha_splash.png"),
    ("Diluc", "diluc_gacha_splash.png"),
    ("Diona", "diona_gacha_splash.png"),
    ("Eula", "eula_gacha_splash.png"),
    ("Fischl", "fischl_gacha_splash.png"),
    ("Ganyu", "ganyu_gacha_splash.png"),
    ("Gorou", "gorou_gacha_splash.png"),
    ("Hu Tao", "hutao_gacha_splash.png"),
    ("Jean", "jean_gacha_splash.png"),
    ("Kaedehara Kazuha", "kazuha_gacha_splash.png"),
    ("Kaeya", "kaeya_gacha_splash.png"),
    ("Kamisato Ayaka", "ayaka_gacha_splash.png"),
    ("Keqing", "keqing_gacha_splash.png"),
    ("Klee", "klee_gacha_splash.png"),
    ("Kujou Sara", "sara_gacha_splash.png"),
    ("Lisa", "lisa_gacha_splash.png"),
    ("Mona", "mona_gacha_splash.png"),
    ("Ningguang", "ningguang_gacha_splash.png"),
    ("Noelle", "noelle_gacha_splash.png"),
    ("Qiqi", "qiqi_gacha_splash.png"),
    ("Raiden Shogun", "shougun_gacha_splash.png"),
    ("Razor", "razor_gacha_splash.png"),
    ("Rosaria", "rosaria_gacha_splash.png"),
    ("Sangonomiya Kokomi", "kokomi_gacha_splash.png"),
    ("Sayu", "sayu_gacha_splash.png"),
    ("Shenhe", "shenhe_gacha_splash.png"),
    ("Sucrose", "sucrose_gacha_splash.png"),
    ("Tartaglia", "tartaglia_gacha_splash.png"),
    ("Thoma", "thoma_gacha_splash.png"),
    ("Venti", "venti_gacha_splash.png"),
    ("Xiangling", "xiangling_gacha_splash.png"),
    ("Xiao", "xiao_gacha_splash.png"),
    ("Xingqiu", "xingqiu_gacha_splash.png"),
    ("Xinyan", "xinyan_gacha_splash.png"),
    ("Yanfei", "feiyan_gacha_splash.png"),
    ("Yoimiya", "yoimiya_gacha_splash.png"),
    ("Yun Jin", "yunjin_gacha_splash.png"),
    ("Zhongli", "zhongli_gacha_splash.png"),
]


root = './img/'
output = './out/'
nclusters = 10
cwidth = 1821 # taken from max
cheight = 50
cshape = (cheight, cwidth, 4)
afilter = 200 # 0 - 255, alpha filter


def load(cvtype):
    titles = []
    files = []
    images = []
    max = [0, 0]
    for data in image_data:
        image = cv2.imread(root + data[1], cvtype)
        titles.append(data[0])
        files.append(data[1])
        images.append(image)
        if image.shape[0] > max[0]:
            max[0] = image.shape[0]
        if image.shape[1] > max[1]:
            max[1] = image.shape[1]

    return titles, files, images, max


def pad_images(images, max):
    result = []
    for image in images:
        y = max[0] - image.shape[0]
        x = max[1] - image.shape[1]
        t = y // 2
        b = y - t
        l = x // 2
        r = x - l
        result.append(cv2.copyMakeBorder(image, t, b, l, r, cv2.BORDER_CONSTANT, value=[0, 0, 0, 0]))

    return result


def rgba_to_rgb(image):
    B, G, R, A = cv2.split(image)
    alpha = A / 255

    R = (255 * (1 - alpha) + R * alpha).astype(np.uint8)
    G = (255 * (1 - alpha) + G * alpha).astype(np.uint8)
    B = (255 * (1 - alpha) + B * alpha).astype(np.uint8)

    return cv2.merge((B, G, R))


def colors_buffer(image):
    # Reshape to a single array
    image = image.reshape((image.shape[0] * image.shape[1], 4))
    size = image.shape[0]

    # Filter pixels that are too transparent
    image = np.array([x for x in image if x[3] > afilter])

    ratio = image.shape[0] / size

    # Make cluster
    cluster = KMeans(n_clusters=nclusters).fit(image)
    centroids = cluster.cluster_centers_

    # Get the number of different clusters, create histogram, and normalize
    labels = np.arange(0, len(np.unique(cluster.labels_)) + 1)
    (hist, _) = np.histogram(cluster.labels_, bins = labels)
    hist = hist.astype("float")
    hist /= hist.sum()

    # Create frequency rect and iterate through each cluster's color and percentage
    freq_rect = np.zeros(cshape, dtype=np.uint8)
    colors = sorted([(percent, color) for (percent, color) in zip(hist, centroids)])
    start = 0
    for (percent, color) in colors:
        end = start + (percent * cwidth)
        cv2.rectangle(freq_rect, (int(start), 0), (int(end), cheight), color.astype("uint8").tolist(), -1)
        start = end

    # Calculate average color of the entire image
    avg_color = np.average(image, axis=0)
    avg = np.full(cshape, avg_color, dtype=np.uint8)

    return freq_rect, avg, ratio


def merge_images(image, freq, avg):
    shape = (image.shape[0] + freq.shape[0] + avg.shape[0], image.shape[1], image.shape[2])
    result = np.full(shape, [0, 0, 0, 1], dtype=np.uint8)

    freq_rgba = cv2.cvtColor(freq, cv2.COLOR_RGB2RGBA)
    avg_rgba = cv2.cvtColor(avg, cv2.COLOR_RGB2RGBA)

    result[0:image.shape[0], 0:image.shape[1], :] = image
    result[image.shape[0]:image.shape[0] + cheight, :, :] = freq_rgba
    result[image.shape[0] + freq.shape[0]:, :, :] = avg_rgba

    return result


if __name__ == '__main__':

    titles, files, images, vmax = load(cv2.IMREAD_UNCHANGED)

    images = pad_images(images, vmax)

    for i in range(len(images)):
        image = images[i]
        freq, avg, ratio = colors_buffer(image)
        result = merge_images(image, freq, avg)
        print(f"{titles[i]:>20} {ratio * 100:>6.2f}% {result.shape}")
        cv2.imwrite(output + titles[i] + '.png', result)
