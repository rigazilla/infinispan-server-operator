package org.infinispan;

import java.security.cert.X509Certificate;

import java.net.Socket;
import javax.net.ssl.HttpsURLConnection;
import javax.net.ssl.KeyManager;
import javax.net.ssl.SSLContext;
import javax.net.ssl.SSLSession;
import javax.net.ssl.TrustManager;
import javax.net.ssl.X509TrustManager;
import javax.net.ssl.X509ExtendedTrustManager;
import java.security.cert.CertificateException;
import javax.net.ssl.SSLEngine;

import org.infinispan.client.hotrod.RemoteCache;
import org.infinispan.client.hotrod.RemoteCacheManager;
import org.infinispan.client.hotrod.configuration.Configuration;
import org.infinispan.client.hotrod.configuration.ConfigurationBuilder;
import org.infinispan.client.hotrod.configuration.ClientIntelligence;
import org.infinispan.client.hotrod.configuration.SaslQop;

import org.infinispan.client.hotrod.impl.ConfigurationProperties;

import gnu.getopt.Getopt;

/**
 * Infinispan+operator+hotrod client quickstart
 *
 */
public class App {
    static String host = "127.0.0.1";
    static int port = ConfigurationProperties.DEFAULT_HOTROD_PORT;
    static String username, password;
    static boolean tls = true;
    static String truststorePath;
    static String truststorePass;

    static SSLContext sslContext;
    public static void main(String[] args) {
        // Accept untrusted certificates for this tutorial
        initSSLSocketFactory();
        // getopt and configure static variables
        configure(args);
        // Create a configuration for a locally-running server

        Configuration configuration;

        if (tls) { // Configuration with encryption
            ConfigurationBuilder cb = new ConfigurationBuilder();
            cb.security().ssl().sslContext(sslContext).enable();
            configuration = cb.addServer().host(host).port(port)
                    .clientIntelligence(ClientIntelligence.BASIC).security().authentication().enable()
                    .username("developer").password(password).realm("default").serverName("infinispan")
                    .saslMechanism("DIGEST-MD5").saslQop(SaslQop.AUTH).ssl()
                    .build();
        } else {
            configuration = new ConfigurationBuilder().addServer().host(host).port(port)
            .clientIntelligence(ClientIntelligence.BASIC).security().authentication().saslMechanism("PLAIN").username(username).password(password).build();
        }
        // Connect to the server
        RemoteCacheManager cacheManager = new RemoteCacheManager(configuration);
        // Obtain the remote cache
        try {
            cacheManager.administration().createCache("quickstart-cache", "org.infinispan.DIST_SYNC");
        } catch (Exception ex) {
            // Maybe cache already exist? Go on anyway
        }
        RemoteCache<String, String> cache = cacheManager.getCache("quickstart-cache");
        /// Store a value
        cache.put("key", "value");
        // Retrieve the value and print it out
        System.out.printf("key = %s\n", cache.get("key"));
        // Stop the cache manager and release all resources
        cacheManager.administration().removeCache("quickstart-cache");
        cacheManager.stop();
    }

    static void configure(String args[]) {
        Getopt g = new Getopt("testprog", args, "h:p:U:P:t:k:");
        //
        int c;
        String arg;
        while ((c = g.getopt()) != -1) {
            switch (c) {
            case 'h':
                arg = g.getOptarg();
                host = arg;
                break;
            case 'p':
                arg = g.getOptarg();
                port = Integer.parseInt(arg);
                break;
            case 'U':
                arg = g.getOptarg();
                username = arg;
                break;
            case 'P':
                arg = g.getOptarg();
                password = arg;
                break;
            case 't':
                tls = true;
                arg = g.getOptarg();
                truststorePath = arg;
                break;
            case 'k':
                tls = true;
                arg = g.getOptarg();
                truststorePass = arg;
                break;
            case '?':
            default:
                break;
            }
        }
    }

    // Create a trust manager that does not validate certificate chains
    private static TrustManager[] _trustAllCerts = new TrustManager[] { new ZeroSecurityTrustManager() };

    private static void initSSLSocketFactory() {

        try {
            sslContext = SSLContext.getInstance("TLS");
            sslContext.init(null, _trustAllCerts, null);
        } catch (Exception e) {
        }
    }
}

class ZeroSecurityTrustManager extends X509ExtendedTrustManager {
    @Override
    public void checkClientTrusted(X509Certificate[] chain, String authType, Socket socket) throws CertificateException {
    }
 
    @Override
    public void checkServerTrusted(X509Certificate[] chain, String authType, Socket socket) throws CertificateException {
 
    }
 
    @Override
    public void checkClientTrusted(X509Certificate[] chain, String authType, SSLEngine engine) throws CertificateException {
 
    }
 
    @Override
    public void checkServerTrusted(X509Certificate[] chain, String authType, SSLEngine engine) throws CertificateException {
 
    }
 
    @Override
    public void checkClientTrusted(X509Certificate[] chain, String authType) throws CertificateException {
 
    }
 
    @Override
    public void checkServerTrusted(X509Certificate[] chain, String authType) throws CertificateException {
 
    }
 
    @Override
    public X509Certificate[] getAcceptedIssuers() {
       return null;
    }
}
