use strict;
use warnings;


my @start = (
	"",
	"DBJV",
	"PVBWRDF",
	"RGFLDCWQ",
	"WJPMLNDB",
	"HNBPCSQ",
	"RDBSNG",
	"ZBPMQFSH",
	"WLF",
	"SVFMR"
);

foreach (@start) {
    print "$_\n";
}

my $sum = 0;
while (my $line = <>) {
    chomp $line;
    my ($cnt,$src,$dst) = $line =~ m/\d+/g;

    print "$line\n";
    $start[$dst] .= reverse(substr($start[$src],-$cnt));
    substr($start[$src],-$cnt) = "";

    print "NOW:\n";
    foreach (@start) {
        print "$_\n";
    }
}

