# Integración de paypal

Proyecto con concepto de Api Restful que integra Paypal como método de pago, backend desarrollado con GOLANG y Frontend con React



## Librerías de terceros

Instalación de driver de postgres
`go get github.com/lib/pq`

Instalación de uuid para emplear caracteres de indentificador
`go get github.com/google/uuid`

Instalación de echo router
`go get github.com/labstack/echo/v4`

Instalación de lector de variables de entorno
`go get github.com/joho/godotenv`



## Configuración de paypal

### Accounts

Seguir los pasos siguientes:

1. Verificar las cuentas existentes (personal/negocio)
2. Podemos crear una nueva cuenta para tener más accesible los datos

### My app and credentials

Contemplar los siguientes puntos:

1. Trabajar en modo sandbox
2. Podemos emplear la aplicacion por defecto (opcional)
3. Podemos crear una nueva cuenta
   - Nombre de la app <!--pasarelapago-->
   - Tipo de la app <!--merchant: pagos directos | platform: pagos desde otra persona-->
   - Correo principal del negocio <!--sb-rroex21951342@business.example.com -->
4. Obtenermos los datos de la app [Sandbox Account|Client ID|Secret]

### Creación del webhook

1. "My app and credentials" > "app a la cual se agregará webhook"
2. Agregamos una URL con complejidad para evitar ataques <!--https://test.com/v1/paypal-->
3. Seleccionamos los eventos "payment sale completed" y "payment capture completed", guardamos.
4. Guardamos el ID y la URL (El ID es necesario para validar que la petición es de paypal) <!--ID:58W05339GE8324617-->



## SSL para proyecto

### Aplicacion con SSL

Eejecutar el comando desde la raíz del proyecto 
`openssl req -x509 -newkey rsa:4096 -keyout key.pem -nodes -out cert.pem -days 365`

Colocamos los datos que nos solicita

Cambiamos Start() por StartTLS() para arrancar con SSL y pasamos las direcciones de los certificados

### Configurar el servidor

Validar que snap está actualizado
`sudo snap install core; sudo snap refresh core`

Instalar CERTBOT
`sudo snap install --classic certbot`

Preparamos CERTBOT con link
`sudo ln -s /snap/bin/certbot /usr/bin/certbot`

Ejecutamos el certbot para generar certificado
`sudo certbot certonly --standalone`

Colocamos el nombre de dominio y los datos que nos solicita
{

IMPORTANT NOTES:
 - Congratulations! Your certificate and chain have been saved at:
   /etc/letsencrypt/live/paypal.cobranza-legal.us/fullchain.pem
   Your key file has been saved at:
   /etc/letsencrypt/live/paypal.cobranza-legal.us/privkey.pem
   Your cert will expire on 2023-01-22. To obtain a new or tweaked
   version of this certificate in the future, simply run certbot again
   with the "certonly" option. To non-interactively renew *all* of
   your certificates, run "certbot renew"
 - If you like Certbot, please consider supporting our work by:

   Donating to ISRG / Let's Encrypt:   https://letsencrypt.org/donate
   Donating to EFF:                    https://eff.org/donate-le






server {

        root /var/www/paypal.cobranza-legal.us;
        index index.html index.htm index.nginx-debian.html;

        server_name paypal.cobranza-legal.us;

        location / {
                try_files $uri $uri/ =404;
        }

    listen [::]:443 ssl ipv6only=on; # managed by Certbot
    listen 443 ssl; # managed by Certbot
    ssl_certificate /etc/letsencrypt/live/paypal.cobranza-legal.us/fullchain.pem; # managed by Certbot
    ssl_certificate_key /etc/letsencrypt/live/paypal.cobranza-legal.us/privkey.pem; # managed by Certbot
    include /etc/letsencrypt/options-ssl-nginx.conf; # managed by Certbot
    ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem; # managed by Certbot

}


}

## Puesta en marcha del proyecto

Configurar auto run de la app
`cd /etc/systemd/system`
`vim paypal.service`

Colocamos el siguiente código 
{
[Unit]
Description=Servicio de paypal
After=network.target

[Service]
Type=simple
User=root
WorkingDirectory=/var/www/pasarelapago
ExecStart=/var/www/pasarelapago/pasarelapago
Restart=always

[Install]
WantedBy=multi-user.target
}

Registramos el servicio
`sudo systemctl enable paypal.service`

Verificamos el stado del servicio
`sudo systemctl status paypal.service`

Arrancamos el servicio servicio
`sudo systemctl start paypal.service`


Hacer el build del proyecto
`go build -o pasarelapago ./cmd`

Ejecutar el proyecto compilado
`./pasarelapago`




server {
   listen 80;
   listen [::]:80;

   root /var/www/your_domain/html;
   index index.html index.htm index.nginx-debian.html;

   server_name your_domain www.your_domain;

   location / {
            try_files $uri $uri/ =404;
   }
}