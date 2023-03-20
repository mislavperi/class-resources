from selenium import webdriver
from selenium.webdriver.chrome.service import Service
from selenium.webdriver.chrome.options import Options
from selenium.webdriver.common.keys import Keys
from selenium.webdriver.common.by import By

sel = Service("./chromedriver")
chrome_options = Options()
chrome_options.add_experimental_option("detach", True)
driver = webdriver.Chrome(service=sel, options=chrome_options)
driver.implicitly_wait(2)
driver.get('https://www.google.com/search?q=algebra&sxsrf=AJOqlzWrBxUo8-MnUQ9FRMqUpd50aTbTTA%3A1678535111859&source=hp&ei=x2kMZKv3KsSeptQP7tG6yAo&iflsig=AK50M_UAAAAAZAx311tqodJFnaUE2DiynLCgc9K-iGUl&ved=0ahUKEwjr9s3I5tP9AhVEj4kEHe6oDqkQ4dUDCAg&uact=5&oq=algebra&gs_lcp=Cgdnd3Mtd2l6EAMyBAgjECcyBAgAEEMyBAgAEEMyCggAEIAEEBQQhwIyCwguEK8BEMcBEIAEMgUIABCABDIFCAAQgAQyBQgAEIAEMgUIABCABDIICAAQgAQQywE6BAguECc6CgguEMcBENEDEEM6CwguEIAEEMcBENEDOgUILhCABDoICC4QgAQQ1AI6BQgAEJECOgQILhBDUABY7ANgqgRoAHAAeACAAccBiAHPBpIBAzAuNpgBAKABAQ&sclient=gws-wiz')
m = driver.find_element(by=By.NAME, value="Algebra: Naslovnica")
print(m)