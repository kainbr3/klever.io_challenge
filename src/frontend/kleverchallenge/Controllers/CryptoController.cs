using Grpc.Net.Client;
using kleverchallenge.Models;
using KleverGrpcClient;
using Microsoft.AspNetCore.Mvc;
using static KleverGrpcClient.CryptoService;
using System.Text.Json;
namespace kleverchallenge.Controllers;

public class CryptoController : Controller
{
    private static GrpcChannel channel = GrpcChannel.ForAddress("http://localhost:50051");
    private CryptoServiceClient client = new CryptoServiceClient(channel);

    [HttpGet]
    public IActionResult List()
    {
        List<Crypto> cryptos = FindAllCryptos().Result.Cryptos.ToList();
        ViewBag.CryptoList = cryptos;

        return View();
    }

    public IActionResult Add()
    {
        //ViewBag.Crypto = AddNewCrypto().Result;
        ViewBag.Crypto = "Test";
        return View();
    }

    public IActionResult Update()
    {
        ViewBag.PageName = "HomePage";
        return View();
    }

    public IActionResult Delete()
    {
        return View();
    }

    [HttpPost]
    public JsonResult Test(string name, string token)
    {
        var crypto = new CryptoEntity
        {
            Id = 999,
            Name = name,
            Token = token,
            Votes = 22,
            Image = "1011"
        };
        return Json(JsonSerializer.Serialize(crypto));
    }
    public async Task<CreateNewCryptoResponse> AddNewCrypto()
    {
        var request = new CreateNewCryptoRequest { Name = "Lascado", Token = "LASC" };
        return await client.CreateNewCryptoAsync(request);
    }

    public async Task<ListCryptosResponse> FindAllCryptos()
    {
        var request = new ListCryptosRequest();
        return await client.ListCryptosAsync(request);
    }
}