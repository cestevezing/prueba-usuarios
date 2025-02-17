# Prueba API Usuarios

## Descripción

Este es un ejercicio para la prueba técnica en la empresa **Tribal Worldwide**. La prueba consiste en consultar una API de usuarios que tiene la capacidad de retornar como máximo **5,000 usuarios diferentes**. El objetivo es desarrollar una API que devuelva **15,000 usuarios**, incluyendo los siguientes campos:

- Género
- Nombre
- Apellido
- Email
- País
- Ciudad
- UUID (valor único)

La API debe responder en menos de **2.5 segundos**.

---

## Análisis del ejercicio 

1. **Tiempo de respuesta de la API original:**

   - La solicitud de **5,000 usuarios** toma entre **6 y 8 segundos**.
   - La solicitud de **100 usuarios** toma en promedio **180 - 210 milisegundos**.

2. **Uso de concurrencia en Go:**

   - Se pueden utilizar **go routines** para ejecutar múltiples solicitudes en paralelo.
   - Sin embargo, se detectó que después de **4 conexiones simultáneas**, el servidor responde con **HTTP 429 (Too Many Requests)**.
   - Por esta razón, la lógica limita el número de conexiones concurrentes a **4**.

3. **Optimización mediante filtros:**

   - La API de generación de usuarios permite aplicar filtros para reducir la cantidad de datos retornados.
   - Se optó por solicitar solo los datos estrictamente necesarios para mejorar el rendimiento.

4. **Estrategia para optimizar el tiempo de respuesta:**

   - Se solicita **un único usuario** con todos los campos requeridos y se usa como plantilla para los **15,000 usuarios**.
   - Luego, se realizan **15,000 solicitudes individuales** para obtener **UUIDs únicos**.

---

## Resultados

- El mejor tiempo logrado para la generación de **15,000 usuarios** fue de **6.3 segundos**.
- El tiempo de respuesta está limitado por las restricciones del servidor de generación de usuarios.

---

## Tecnologías utilizadas

- **Golang** (concurrencia con goroutines)
- **Fiber** para la construcción del api

---

## Instalación y ejecución

1. Clonar el repositorio:

   ```sh
   git clone https://github.com/cestevezing/prueba-usuarios.git
   cd prueba-usuarios
   ```

2. Instalar dependencias:

   ```sh
   go mod tidy
   ```

3. Ejecutar la API:

   ```sh
   go run main.go
   ```

---

## Contacto

Aunque el resultado no fue el esperado para esta prueba quise subir el proyecto para recibir feedback con la solución a este problema, ya que fue un reto el poder trabajar con este problema y me gustaría aprender de la solución.

Muchas gracias por la oportunidad, quedo atento a sus comentarios


