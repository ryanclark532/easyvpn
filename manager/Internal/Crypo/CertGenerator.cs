using Org.BouncyCastle.Asn1.Sec;
using Org.BouncyCastle.Asn1.X509;
using Org.BouncyCastle.Crypto.Operators;
using Org.BouncyCastle.Crypto.Parameters;
using Org.BouncyCastle.Crypto.Prng;
using Org.BouncyCastle.Math;
using Org.BouncyCastle.Security;
using Org.BouncyCastle.X509;

namespace manager.Internal.Crypo;

public class CertGenerator(string outputPath)
{
    public void GenerateSignedCertificate(string name)
    {
        var random = new SecureRandom(new CryptoApiRandomGenerator());

        var root = Crypto.LoadCertificateFromPem(File.ReadAllText(Path.Combine(outputPath, "root.crt")));
        var rootKey = Crypto.LoadPrivateKeyFromPem(File.ReadAllText(Path.Combine(outputPath, "root.key")));

        var keyPairGenerator = GeneratorUtilities.GetKeyPairGenerator("ECDSA");
        keyPairGenerator.Init(new ECKeyGenerationParameters(SecObjectIdentifiers.SecP384r1, random));
        var keyPair = keyPairGenerator.GenerateKeyPair();

        var generator = new X509V3CertificateGenerator();
        var serialNumber = BigInteger.ProbablePrime(128, random);
        generator.SetSerialNumber(serialNumber);
        generator.SetIssuerDN(root.SubjectDN);
        generator.SetSubjectDN(new X509Name($"CN={name}, O=Example Organization"));
        generator.SetNotBefore(DateTime.UtcNow);
        generator.SetNotAfter(DateTime.UtcNow.AddYears(1));
        generator.SetPublicKey(keyPair.Public);

        var signatureFactory = new Asn1SignatureFactory("SHA384withECDSA", rootKey, random);
        var signedCert = generator.Generate(signatureFactory);

        File.WriteAllText(Path.Combine(outputPath, $"{name}.crt"), Crypto.ExportToPem(signedCert));
        File.WriteAllText(Path.Combine(outputPath, $"{name}.key"), Crypto.ExportToPem(keyPair.Private));
    }}