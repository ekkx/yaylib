import unittest

from yaylib.config import VERSION_NAME
from yaylib.device import DEVICES, Device

TEST_DEVICE_TYPE = "android"
TEST_OS_VERSION = "11"
TEST_SCREEN_DENSITY = "3.5"
TEST_SCREEN_SIZE = "1440x2960"
TEST_MODEL = "Galaxy S9"


class TestDevice(unittest.TestCase):
    def test_create_device_with_model(self):
        created_device = Device.create(model=TEST_MODEL)
        self.assertIsNotNone(created_device)
        self.assertEqual(created_device.model, TEST_MODEL)
        self.assertEqual(created_device.device_type, TEST_DEVICE_TYPE)
        self.assertEqual(created_device.os_version, TEST_OS_VERSION)
        self.assertEqual(created_device.screen_density, TEST_SCREEN_DENSITY)
        self.assertEqual(created_device.screen_size, TEST_SCREEN_SIZE)

    def test_create_random_device(self):
        created_device = Device.create()
        created_device_dict = next(
            (device for device in DEVICES if device["model"] == created_device.model),
            None,
        )
        self.assertIsNotNone(created_device)

        created_device_type = created_device_dict.get("device_type")
        created_os_version = created_device_dict.get("os_version")
        created_screen_density = created_device_dict.get("screen_density")
        created_screen_size = created_device_dict.get("screen_size")

        self.assertEqual(created_device.device_type, created_device_type)
        self.assertEqual(created_device.os_version, created_os_version)
        self.assertEqual(created_device.screen_density, created_screen_density)
        self.assertEqual(created_device.screen_size, created_screen_size)

    def test_device_user_agent(self):
        excpected_user_agent = f"{TEST_DEVICE_TYPE} {TEST_OS_VERSION} ({TEST_SCREEN_DENSITY}x {TEST_SCREEN_SIZE} {TEST_MODEL})"
        created_device = Device.create(model=TEST_MODEL)
        self.assertEqual(created_device.get_user_agent(), excpected_user_agent)

    def test_device_info(self):
        excpected_user_agent = f"{TEST_DEVICE_TYPE} {TEST_OS_VERSION} ({TEST_SCREEN_DENSITY}x {TEST_SCREEN_SIZE} {TEST_MODEL})"
        excpected_device_info = f"yay {VERSION_NAME} {excpected_user_agent}"
        created_device = Device.create(model=TEST_MODEL)
        self.assertEqual(created_device.get_device_info(), excpected_device_info)

    def test_create_device_with_invalid_model(self):
        with self.assertRaises(ValueError):
            Device.create(model="__invalid_model__")
