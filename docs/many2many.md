# Many2Many

## What happens if you switch the order of the statements wgp.Wait() and close(ch) in the end of the main function?

### Hypothesis

My hypothesis is that the program won't run since the main routine will close the channel before the goroutines finish sending on it.

### Result

As stated in the hypothesis it didn't run.

## What happens if you move the close(ch) from the main function and instead close the channel in the end of the function Produce?
    
### Hypothesis

For the same reason as stated above it won't run. The first producer will close the channel but the other producers will still try to send on it.

### Result

yup.

## What happens if you remove the statement close(ch) completely?
    
### Hypothesis

The consume goroutines will never end since their loop is waiting on the channel to close. They close anyhow when the main routine ends (although I think it might be a bit more messy closing down) so the program should execute the same way regardless.

### Result

I can't see any visible diffrence in performance. So I'll assume the hypothesis was correct. 

## What happens if you increase the number of consumers from 2 to 4?

### Hypothesis

We have more threads ready to recieve so there should be less downtime where a producer is waiting for a consumer to read it's output. Therefore the program should finish faster

### Result

It finishes faster

## Can you be sure that all strings are printed before the program stops?

### Hypothesis

No you can't. Most strings will be printed because each producer has to wait until a consumer takes it output before producing it's next output. Therefore the program can't finnish until after the last two messages have been recieved. However there is no guarantee that both these messages will have time to print before the main routine ends.

### Result

Seems like this was correct as well. From my testing it sometimes doesn't print one message.