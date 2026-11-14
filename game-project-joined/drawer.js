const toggleDrawer = (id) => {
  const element = document.getElementById(id)

  if (element.classList.contains("hiddenDrawer")) {
    element.classList.remove("hiddenDrawer")
  } else {
    element.classList.add("hiddenDrawer")
  }
}
