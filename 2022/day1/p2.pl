use strict;
use warnings;

my $sum = 0;
my @totals;
while (my $line = <>) {
    chomp;
    if ($line == "") {
        push(@totals, $sum);
        $sum = 0;
    }
    $sum += $line;
}
@totals = sort {$a <=> $b} @totals;
print "@totals\n";

print $totals[-1] + $totals[-2] + $totals[-3], "\n";

