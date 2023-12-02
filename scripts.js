document.addEventListener('DOMContentLoaded', () => {
    const searchInput = document.querySelector('#idInput');
    const searchButton = document.querySelector('#searchButton');
    const updateButton = document.querySelector('#updateButton');
    const deleteButton = document.querySelector('#deleteButton');
    const resultadoContainer = document.querySelector('#resultado');
    const errorContainer = document.querySelector('#error');

    const apiUrl = 'http://localhost:8080/estudiantes/';

    const obtenerDatos = async () => {
        const id = searchInput.value.trim();

        if (!Number.isInteger(Number(id))) {
            mostrarError();
            return;
        }

        try {
            const url = `${apiUrl}${id}`;
            const response = await axios.get(url);

            if (response.status === 200) {
                const estudiante = response.data;
                llenarCasillas(estudiante);
            } else {
                mostrarError();
            }
        } catch (error) {
            mostrarError();
        }
    };

    const actualizarEstudiante = async () => {
        const id = searchInput.value.trim();

        if (!Number.isInteger(Number(id))) {
            mostrarError();
            return;
        }

        const nuevoEstudiante = {
            nombre: document.getElementById('nombre').value,
            edad: document.getElementById('edad').value,
            grado: document.getElementById('grado').value,
            jornada: document.getElementById('jornada').value,
            direccion: document.getElementById('direccion').value,
            telefono: document.getElementById('telefono').value,
            correo: document.getElementById('correo').value,
        };

        try {
            const url = `${apiUrl}${id}`;
            const response = await axios.patch(url, nuevoEstudiante);

            if (response.status === 200) {
            } else {
                mostrarError();
            }
        } catch (error) {
            mostrarError();
        }
    };

    const addButton = document.querySelector('#addButton');
    addButton.addEventListener('click', agregarEstudiante);

    async function agregarEstudiante() {
        const nombre = document.querySelector('#nombre').value.trim();
        const edad = parseInt(document.querySelector('#edad').value.trim(), 10);
        const grado = document.querySelector('#grado').value.trim();
        const jornada = document.querySelector('#jornada').value.trim();
        const direccion = document.querySelector('#direccion').value.trim();
        const telefono = parseInt(document.querySelector('#telefono').value.trim(), 10);
        const correo = document.querySelector('#correo').value.trim();
    
        if (!nombre || isNaN(edad) || !grado || !jornada || !direccion || isNaN(telefono) || !correo) {
            mostrarError();
            return;
        }
    
        try {
            const url = 'http://localhost:8080/estudiantes';
            const response = await axios.post(url, {
                nombre: nombre,
                edad: edad,
                grado: grado,
                jornada: jornada,
                direccion: direccion,
                telefono: telefono,
                correo: correo
            });
    
            if (response.status === 201) {
                const nuevoId = response.data;
                resultadoContainer.style.display = 'none';
                console.log('Nuevo estudiante creado con ID:', nuevoId);
                
            } else {
                mostrarError();
            }
        } catch (error) {
            mostrarError();
        }
    }

    function limpiarFormulario() {
        // Limpiar los valores en las casillas después de agregar un estudiante
        document.querySelector('#nombre').value = '';
        document.querySelector('#edad').value = '';
        document.querySelector('#grado').value = '';
        document.querySelector('#jornada').value = '';
        document.querySelector('#direccion').value = '';
        document.querySelector('#telefono').value = '';
        document.querySelector('#correo').value = '';
    }


    const eliminarEstudiante = async () => {
        const id = searchInput.value.trim();

        if (!Number.isInteger(Number(id))) {
            mostrarError();
            return;
        }

        try {
            const url = `${apiUrl}${id}`;
            const response = await axios.delete(url);

            if (response.status === 200) {
                limpiarFormulario()
                
            } else {
                mostrarError();
            }
        } catch (error) {
            mostrarError();
        }
    };

    const llenarCasillas = (estudiante) => {
        document.getElementById('nombre').value = estudiante.nombre;
        document.getElementById('edad').value = estudiante.edad;
        document.getElementById('grado').value = estudiante.grado;
        document.getElementById('jornada').value = estudiante.jornada;
        document.getElementById('direccion').value = estudiante.direccion;
        document.getElementById('telefono').value = estudiante.telefono;
        document.getElementById('correo').value = estudiante.correo;

        resultadoContainer.style.display = 'block';
        errorContainer.style.display = 'none';
    };

    const mostrarError = () => {
        errorContainer.innerHTML = '<p>¡Ups! Revisa la información suministrada.</p>';
        errorContainer.style.display = 'block';
        resultadoContainer.style.display = 'none';
    };

    // Asignar funciones a los eventos de los botones
    searchButton.addEventListener('click', obtenerDatos);
    updateButton.addEventListener('click', actualizarEstudiante);
    deleteButton.addEventListener('click', eliminarEstudiante);

});



