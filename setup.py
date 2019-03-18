from setuptools import find_packages, setup
from distutils.util import convert_path


def get_version(pkg_name):
    ver_path = convert_path('{}/version.py'.format(pkg_name))
    with open(ver_path) as ver_file:
        text = ver_file.read()
    try:
        ver = text.split('version = ')[1]
    except IndexError:
        raise RuntimeError('Could not parse version string: {}'.format(text))
    return ver.strip('"').strip("'")


pkg_name = 'vprint'
setup(
    name='vprint',
    version=get_version('vprint'),
    description='Verbose Printing Utility',
    packages=find_packages(),

)
