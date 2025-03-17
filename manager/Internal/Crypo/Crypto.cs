using Org.BouncyCastle.Crypto;
using Org.BouncyCastle.X509;
using Org.BouncyCastle.OpenSsl;


public static class Crypto
{
    
    public static string ExportToPem(object obj)
    {
        using var sw = new StringWriter();
        
        var pemWriter = new PemWriter(sw);
        pemWriter.WriteObject(obj);
        return sw.ToString();
    }

    public static X509Certificate LoadCertificateFromPem(string pem)
    {
        using var sr = new StringReader(pem);
        var pemReader = new PemReader(sr);
        return (X509Certificate)pemReader.ReadObject();
    }

    public static AsymmetricKeyParameter LoadPrivateKeyFromPem(string pem)
    {
        using var sr = new StringReader(pem);
        var pemReader = new PemReader(sr);
        return ((AsymmetricCipherKeyPair)pemReader.ReadObject()).Private;
    }
    
}