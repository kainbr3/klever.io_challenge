using Microsoft.AspNetCore.Mvc;
namespace kleverchallenge.Controllers;

public class HomeController : Controller
{
    public IActionResult Index()
    {
        return View();
    }
}