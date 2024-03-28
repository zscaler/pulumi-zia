"""A Python Pulumi program"""

import zscaler_pulumi_zia as zia

example = zia.ZIAURLCategories("zia-url-category-example",
    super_category="USER_DEFINED",
    configured_name="Pulumi ZIA URL Category",
    description="Pulumi ZIA URL Category",
    keywords=["microsoft"],
    custom_category=True,
    type="URL_CATEGORY",
    urls=[
        ".coupons.com",
        ".resource.alaskaair.net",
        ".techrepublic.com",
        ".dailymotion.com",
        ".osiriscomm.com",
        ".uefa.com",
        ".Logz.io",
        ".alexa.com",
        ".baidu.com",
        ".cnn.com",
        ".level3.com",
    ])