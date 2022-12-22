use strict;
use warnings;

my @dirs;
my %size;
while (my $l = <>) {
    if ($l =~ m/ cd \.\./) {
        pop @dirs;
    } elsif ($l =~ m/ cd (.*)/) {
        push @dirs, $1;
        # print $l, "\n";
        # print "@dirs\n";
    } elsif ($l =~ m/^dir /) {
    } elsif ($l =~ m/^(\d+)/) {
        my @allDs = @dirs;
        while(@allDs) {
            $size{join('/',@allDs)} += $1;
            pop(@allDs)
        }
        $size{''} += $1;
    }
}

my $sum_small;
foreach my $k (keys %size) {
    print "/$k: $size{$k}\n";

    $sum_small += $size{$k} if $size{$k} < 100000;
}
print "SUM SMALL: $sum_small\n";

my $total = 70000000;
my $goal = 30000000;

my $needed = $size{''} + $goal - $total;
my $min = $total;
foreach my $v (values %size) {
    $min = $v if $v < $min && $v >= $needed;
}
print "$min\n";
