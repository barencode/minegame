1. Regular function call leads to deadlock fatal because receiver never reached and never gets record in regular function call. worker() stuck in sending, waiting for receiver.
2. Run 'worker()' as 'go worker()' make this call dedicated routine and program flow gets successfully to receivers, that are below that call.
3. WRONG: main() ------worker()----X main
   RIGHT: main() --               -- main
                   |_go worker()_|