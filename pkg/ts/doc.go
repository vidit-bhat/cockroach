// Copyright 2016 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

/*
Package ts provides a basic time series database on top of the underlying
CockroachDB key/value datastore. It is used to server basic metrics generated by
CockroachDB.

An alternative, more comprehensive to this package documentation is available
as a technical note (docs/tech-notes/timeseries.md).

Storing time series data is a unique challenge for databases. Time series data
is typically generated at an extremely high volume, and is queried by providing
a range of time of arbitrary size, which can lead to an enormous amount of data
being scanned for a query. Many specialized time series databases already exist
to meet these challenges; those solutions are built on top of specialized
storage engines which are often unsuitable for general data storage needs, but
currently superior to CockroachDB for the purpose of time series data.

However, it is a broad goal of CockroachDB to provide a good experience for
developers, and out-of-the-box recording of internal metrics has proven to be a
good step towards that goal. This package provides a specialized time series
database, relatively narrow in scope, that can store this data with good
performance characteristics.

# Organization Structure

Time series data is organized on disk according to two basic, sortable properties:
+ Time series name (i.e "sql.operations.selects")
+ Timestamp

This is optimized for querying data for a single series over multiple
timestamps: data for the same series at different timestamps is stored
contiguously.

# Downsampling

The amount of data produced by time series sampling can be considerable; storing
every incoming data point with perfect fidelity can command a tremendous amount
of computing and storage resources.

However, in many use cases perfect fidelity is not necessary; the exact time a
sample was taken is unimportant, with the overall trend of the data over time
being far more important to analysis than the individual samples.

With this in mind, CockroachDB downsamples data before storing it; the original
timestamp for each data point in a series is not recorded. CockroachDB instead
divides time into contiguous slots of uniform length (currently 10 seconds); if
multiple data points for a series fall in the same slot, only the most recent
sample is kept.

In addition to the on-disk downsampling, queries may request further
downsampling for returned data. For example, a query may request one datapoint
be returned for every 10 minute interval, even though the data is stored
internally at a 10 second resolution; the downsampling is performed on the
server side before returning the data. One restriction is that a query cannot
request a downsampling period which is shorter than the smallest on-disk
resolution (e.g. one data point per second).

# Slab Storage

In order to use key space efficiently, we pack data for multiple contiguous
samples into "slab" values, with data for each slab stored in a CockroachDB key.
This is done by again dividing time into contiguous slots, but with a longer
duration; this is known as the "slab duration". For example, CockroachDB
downsamples its internal data at a resolution of 10 seconds, but stores it with
a "slab duration" of 1 hour, meaning that all samples that fall in the same hour
are stored at the same key. This strategy helps reduce the number of keys
scanned during a query.

# Source Keys

Another common use case of time series queries is the aggregation of multiple
series; for example, you may want to query the same metric (e.g. "queries per
second") across multiple machines on a cluster, and aggregate the result.

Specialized Time-series databases can often aggregate across arbitrary series;
however, CockroachDB is specialized for aggregation of the same series across
different machines or disks.

This is done by creating a "source key", typically a node or store ID, which is
an optional identifier that is separate from the series name itself. The source
key is appended to the key as a suffix, after the series name and timestamp;
this means that data that is from the same series and time period, but from
different nodes, will be stored contiguously in the key space. Data from all
sources in a series can thus be queried in a single scan.

# Multiple resolutions

In order to save space on disk, the database stores older data for time series
at lower resolution, more commonly known as a "rollup".

Each single series is recorded initially at a resolution of 10 seconds - one
point is recorded for every ten second interval. Once the data has aged past a
configurable threshold (default 10 days), the data is "rolled" up so that it
has a single datapoint per 30 minute period. This 30-minute resolution data is
retained for 90 days by default.

Note that each rolled-up datapoint contains the first, last, min, max, sum,
count and variance of the original 10 second points used to create the 30
minute point; this means that any downsampler that could have been used on the
original data is still accessible in the rolled-up data.

# Example

A hypothetical example from CockroachDB: we want to record the available
capacity of all stores in the cluster.

The series name is: cockroach.capacity.available

Data points for this series are automatically collected from all stores. When data points are
written, they are recorded with a source key of: [store id]

There are 3 stores which contain data: 1, 2 and 3.  These are arbitrary and may
change over time.

Data is recorded for January 1st, 2016 between 10:05 pm and 11:05 pm. The data
is recorded at a 10 second resolution.

The data is recorded into keys structurally similar to the following:

	tsd.cockroach.capacity.available.10s.403234.1
	tsd.cockroach.capacity.available.10s.403234.2
	tsd.cockroach.capacity.available.10s.403234.3
	tsd.cockroach.capacity.available.10s.403235.1
	tsd.cockroach.capacity.available.10s.403235.2
	tsd.cockroach.capacity.available.10s.403235.3

Data for each source is stored in two keys: one for the 10 pm hour, and one
for the 11pm hour. Each key contains the tsd prefix, the series name, the
resolution (10s), a timestamp representing the hour, and finally the series key. The
keys will appear in the data store in the order shown above.

(Note that the keys will NOT be exactly as pictured above; they will be encoded
in a way that is more efficient, but is not readily human readable.)
*/
package ts
