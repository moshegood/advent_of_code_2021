use strict;
use warnings;

my $sum = 0;
my $max = 0;
while (my $line = <>) {
    chomp;
    if ($line == "") {
        $max = $max > $sum ? $max : $sum;
        $sum = 0;
    }
    $sum += $line;
}
print "$max\n";
