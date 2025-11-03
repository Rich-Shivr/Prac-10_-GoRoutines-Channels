package main

import (
	"fmt"
	"sync" //Paquete que permite coordinar la ejecucion de las goroutines
	"time"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) { //El sync.Waitgroup es para esperar a que un grupo de gouroutines terminen su tarea, para ahora si continuar
	defer wg.Done() // Se ejecuta cuando algun worker ha terminado sus tareas, restando uno
	for job := range jobs {
		fmt.Printf("Worker %d comenzó tarea %d\n", id, job)
		time.Sleep(time.Second) // simula trabajo
		fmt.Printf("Worker %d terminó tarea %d\n", id, job)
	}
	fmt.Printf("Worker %d terminó todas sus tareas\n", id)
}

func main() {
	const numJobs = 20   //Aqui indica cuantas tareas se van a hacer
	const numWorkers = 4 // Aqui indica cuantos trabajadores (workers) van a llevar a cabo las tareas.

	jobs := make(chan int, numJobs) // Crea un channel de tipo entero con buffer, es decir, que enviará las tareas y no se bloquerá, a pesar de que no haya alguien que las recuba, en este caso la cantidad de enteros que envia sera 20, ya que es el valor de numJobs
	var wg sync.WaitGroup           //Aqui se declara el grupo de espera para saber cuando salir del programa.

	// Lanzar los workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1) // Para indicar que hay un worker activo
		go worker(i, jobs, &wg)
	}

	// Enviar tareas
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // Para cerrar el channel.

	// Esperar a que todos los workers terminen
	wg.Wait() // Para esperar hasta que se llamen a todos los wg.Done(), indicando que todos los worker terminaron
	fmt.Println("Todas las tareas han sido procesadas.")
}
