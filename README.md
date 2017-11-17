# limitedcache

httpcache compatible cache with limited size on disk. Size is limited to number of items.
Additionly this cache sends operation types, keys, errors to a channel. This can be used to make changes in cache externally - eg. remove some content earlier from cache.

Code is heavily copied from  github.com/gregjones/httpcache