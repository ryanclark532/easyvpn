using Org.BouncyCastle.Asn1.Sec;
using Org.BouncyCastle.Asn1.X509;
using Org.BouncyCastle.Crypto.Operators;
using Org.BouncyCastle.Crypto.Parameters;
using Org.BouncyCastle.Crypto.Prng;
using Org.BouncyCastle.Math;
using Org.BouncyCastle.Security;
using Org.BouncyCastle.X509;

namespace manager.Internal.Crypo;

public class CaGenerator(string outputDir)
{
    public void GenerateRootCa()
    {
        var random = new SecureRandom(new CryptoApiRandomGenerator());

        var pairGenerator = GeneratorUtilities.GetKeyPairGenerator("ECDSA");
        pairGenerator.Init(new ECKeyGenerationParameters(SecObjectIdentifiers.SecP384r1, random));
        var key = pairGenerator.GenerateKeyPair();

        var generator = new X509V3CertificateGenerator();
        var serialNumber = BigInteger.ProbablePrime(128, random);
        generator.SetSerialNumber(serialNumber);
        generator.SetIssuerDN(new X509Name("CN=Root CA"));
        generator.SetSubjectDN(new X509Name("CN=Root CA"));
        generator.SetNotBefore(DateTime.UtcNow);
        generator.SetNotAfter(DateTime.UtcNow.AddYears(10)); 
        generator.SetPublicKey(key.Public);

        var signatureFactory = new Asn1SignatureFactory("SHA384withECDSA", key.Private, random);

        var cert = generator.Generate(signatureFactory);
        
        File.WriteAllText(Path.Combine(outputDir, "root.crt"), Crypto.ExportToPem(cert));
        File.WriteAllText(Path.Combine(outputDir, "root.key"), Crypto.ExportToPem(key.Private));
    }

}