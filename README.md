# big-o-notation
Big O describes the function growth when the input scales, measuring the algorithm complexity.

# Why is it important to measure performance?
Code is not just "making it work" and it's fine, we follow that rule when we started working with code but unfortunately,
that's not how it works in the real world. In big companies that's a subject really important because we know our services
will receive a huge amount of requests and we must be prepared for that. Our code must work and more than that, it must
work in any scenario, it should be created in the best way possible.

There are many solutions to improve performance a code such as caching data, scaling an application using green threads
(go routines) or threads and so on. Each one of these solution is valid to consider when improving a code performance
but they might be isolated to a specific case, for example caching would make sense when the code is getting some data
from a database for example. The Big O notation will measure the logic complexity of an algorithm and how it grows when
the input size increases