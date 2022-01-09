using Grpc.Net.Client;
using KleverGrpcClient;
using Microsoft.AspNetCore.Mvc;
using static KleverGrpcClient.CryptoService;
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
        ViewBag.Crypto = AddNewCrypto().Result;
        //ViewBag.Crypto = "Test";
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

    public async Task<NewCryptoResponse> AddNewCrypto()
    {
        var request = new NewCryptoRequest{ Name = "Lascado", Token = "LASC" };
        return await client.CreateNewCryptoAsync(request);
    }

    public async Task<ListCryptosResponse> FindAllCryptos()
    {
        var request = new ListCryptosRequest();
        return await client.GetCryptosAsync(request);
    }
}