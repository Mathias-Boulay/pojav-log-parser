from enum import Enum


en_locale = {
    'mipmap.enabled': 'Mipmap is not disabled, this may destroy performance',
    'libgles.wrong.value': 'Improper LIBGL_ES value !',
    'libgl.custom.value': 'Custom openGL declaration !',
    'java.args.unparsed': 'Some java arguments were not parsed',
    'java.args.unrecognised': 'Java arguments were not recognised',
    'branch.not.mainline': 'The branch is not mainline',
    'java.runtime.wrong.source': 'Unexpected source for the java runtime !',
    'java.runtime.wrong.type': 'The runtime is not a JRE !',
    'renderer.unsupported.arch': 'Unsupported renderer on 32 bits !',
    'install.forge.corrupted': 'Corrupted forge install detected',
    'load.shader.fail': 'Some shaders failed to be loaded',
    'runtime.early.fail': 'The program failed early',
}


def localize(localizable_array: list[str]) -> list[str]:
    """When the request contains it, display a plain english version of the text"""
    return list(map(lambda x: en_locale.get(x, x), localizable_array))
