namespace kleverchallenge.Models;

public class CryptoViewModel
{
    public int Id { get; set; }
    public string? RequestId { get; set; }

    public bool ShowRequestId => !string.IsNullOrEmpty(RequestId);
}

public class CryptoEntity
{
    public int Id { get; set; }
    public string? Name { get; set; }
    public string? Token { get; set; }
    public int Votes { get; set; }
    public string? Image { get; set; }

    // public CryptoEntity()
    // {
    //     Image = GetCyptoImage(Token);
    // }

    // public string GetCyptoImage(string? tokenName)
    // {
    //     return "";
    // }
}