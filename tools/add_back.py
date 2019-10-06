#! /usr/bin/env python3
import os
import sys
from typing import List


README = "readme.md"
BACK_URL = "[root path](../readme.md)"


# def get_root_path() -> str:
#     curdir = os.path.dirname(__file__)
#     fatherdir = os.path.dirname(curdir)
#     return fatherdir


def has_readme_curdir(path) -> bool:
    """
    当前目录下是否存在 readme 文件
    """

    *_, curdir_files = next(os.walk(path))

    for f in curdir_files:
        if f.lower() == README:
            return True
    return False


def add_back_url(path: str) -> None:
    with open(path, "r") as fr:
        content = fr.readlines()

    def is_need_add(conlist: List) -> bool:
        for line in conlist:
            if line.strip():
                if BACK_URL in line:
                    return False
                else:
                    return True
        return True

    target_with_cr = f"{BACK_URL}\n"
    if is_need_add(content):
        content.insert(0, target_with_cr)

    if is_need_add(content[::-1]):
        list(map(content.append, ("\n", target_with_cr)))

    # 确保开头与结尾都有 back_url
    if content.count(target_with_cr) < 2:
        content.append(target_with_cr)
    # 写入文件
    with open(path, "w") as fw:
        fw.write("".join(content))


def backup_file(path):
    pathbak = f"{path}.bak"
    if os.path.exists(pathbak):
        os.remove(pathbak)
    # WARN 一定要 close
    with os.popen(f"cp {path} {pathbak}") as p:
        p.close()


def main():
    if len(sys.argv) < 2:
        path = "."

    if not has_readme_curdir("."):
        print("不存在 readme 文件")

    path = f"./{README}"
    backup_file(path)
    add_back_url(path)


if __name__ == "__main__":
    main()
