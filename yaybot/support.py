import huepy


def console_print(text, color=None):
    text = '\n' + text
    if color is not None:
        text = getattr(huepy, color)(text)
    print(text)
