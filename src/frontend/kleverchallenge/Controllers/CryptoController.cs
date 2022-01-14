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

    public IActionResult Manage()
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

    public IActionResult Update(int cryptoId)
    {
        ViewBag.PageName = "HomePage";
        CryptoEntity model = new CryptoEntity();
        var result = FindAllCryptos().Result.Cryptos.ToList().Where(x => x.Id == cryptoId).FirstOrDefault();

        if(result != null)
        {
            model.Id = result.Id;
            model.Name = result.Name;
            model.Token = result.Token;
            model.Votes = result.Votes;
        }
        
        return View(model);
    }

    public async Task<IActionResult> Subscribe(int cryptoId)
    {
        var model = new CryptoEntity();
        var obj = new ObserveCryptoRequest{Id = cryptoId};
        ObserveCryptoResponse result;

        using (var response = client.ObserveCrypto(obj))
        {
            while (await response.ResponseStream.MoveNext(CancellationToken.None))
            {
                result = response.ResponseStream.Current;
                model.Id = result.Crypto.Id;
                model.Name = result.Crypto.Name;
                model.Votes = result.Crypto.Votes;
            }
        }

        return View(model);
    }

    [HttpPost]
    public async Task<JsonResult> Insert(string name, string token)
    {
        if(!string.IsNullOrEmpty(name) && !string.IsNullOrEmpty(token)){
            var result = await AddNewCrypto(name, token);
            return Json(result);

        }
        else{
            return Json("Unable to save crypto");
        }
    }

    [HttpPost]
    public async Task<JsonResult> Update([FromBody] CryptoEntity crypto)
    {
        var result = await UpdateCrypto(crypto);
        return Json(result);
    }

    [HttpPost]
    public async Task<JsonResult> Upvote(int cryptoId)
    {
        var result = await UpvoteCrypto(cryptoId);
        return Json(result);
    }   
    
    [HttpPost]
    public async Task<JsonResult> Downvote(int cryptoId)
    {
        var result = await DownvoteCrypto(cryptoId);
        return Json(result);
    }   

    [HttpPost]
    public async Task<JsonResult> Remove(int cryptoId)
    {
        var result = await RemoveCrypto(cryptoId);
        return Json(result);
    } 

    public async Task<CreateNewCryptoResponse> AddNewCrypto(string name, string token)
    {
        var request = new CreateNewCryptoRequest { Name = name, Token = token };
        var result = await client.CreateNewCryptoAsync(request);
        return result;
    }

    public async Task<ListCryptosResponse> FindAllCryptos()
    {
        var request = new ListCryptosRequest();
        return await client.ListCryptosAsync(request);
    }

    public async Task<UpdateCryptoResponse> UpdateCrypto(CryptoEntity crypto)
    {
        var request = new UpdateCryptoRequest{
            Crypto = new Crypto{
                Id = crypto.Id,
                Name = crypto.Name,
                Token = crypto.Token,
                Votes = crypto.Votes
            }
        };

        return await client.UpdateCryptoAsync(request);
    }

    public async Task<EmptyResponse> UpvoteCrypto(int id)
    {
        var request = new UpvoteCryptoRequest{ Id = id};
        return await client.UpvoteCryptoAsync(request);
    }

    public async Task<EmptyResponse> DownvoteCrypto(int id)
    {
        var request = new DownvoteCryptoRequest{ Id = id};
        return await client.DownvoteCryptoAsync(request);
    }

    public async Task<EmptyResponse> RemoveCrypto(int id)
    {
        var request = new DeleteCryptoRequest{ Id = id};
        return await client.DeleteCryptoAsync(request);
    }
}